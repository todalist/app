package main

import (
	"bytes"
	"github.com/goccy/go-yaml"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type GenConfig struct {
	Models      *[]*ModelStruct
	PackageBase string `yaml:"packageBase"`
}

type ModelStruct struct {
	Uname       string `yaml:"-"`
	Name        string
	Fields      *[]*ModelField
	Queriers    []string
	UpdateOmits []string `yaml:"updateOmits"`
	Rewrite     bool
	PackageBase string `yaml:"-"`
}

type ModelField struct {
	Uname     string `yaml:"-"`
	Name      string
	Type      string
	ExtraTags string `yaml:"extraTags"`
}

func main() {
	templateConfigFilePath := os.Getenv("TEMPLATE_CONFIG_PATH")
	basePath := os.Getenv("GEN_CODE_BASE_PATH")
	if templateConfigFilePath == "" {
		panic("template config file path is require. specific via 'TEMPLATE_CONFIG_PATH' env")
	}
	if basePath == "" {
		panic("gen code base path is require. specific via 'GEN_CODE_BASE_PATH' env")
	}
	configContent, err := os.ReadFile(templateConfigFilePath)
	if err != nil {
		panic(err)
	}
	config := new(GenConfig)
	if err = yaml.Unmarshal(configContent, config); err != nil {
		panic(err)
	}

	// fill inner field
	completeModelStruct(config.Models, &config.PackageBase)

	// generate all modules
	genAllModules(basePath, config.Models)
}

func completeModelStruct(models *[]*ModelStruct, packageBase *string) {
	for _, model := range *models {
		model.Uname = strings.ToUpper(model.Name[0:1]) + model.Name[1:]
		model.PackageBase = *packageBase
		for _, field := range *model.Fields {
			field.Uname = strings.ToUpper(field.Name[0:1]) + field.Name[1:]

		}
	}
}

func genAllModules(basePath string, models *[]*ModelStruct) {
	for _, model := range *models {
		genModule(basePath, model)
	}
}

func genModule(basePath string, model *ModelStruct) {
	// check directory is exists
	modulePath := filepath.Join(basePath, model.Name)
	if _, err := os.Stat(modulePath); os.IsExist(err) {
		if !model.Rewrite {
			println("skip gen module: " + model.Name)
			return
		}
		// delete directory ignore all files
		println("previous module will be deleted: " + model.Name)
		os.RemoveAll(modulePath)
	}
	implModulePath := filepath.Join(modulePath, "impl")
	err := os.MkdirAll(implModulePath, 0755)
	if err != nil {
		panic(err)
	}

	// model file
	RenderAndSave(model, MODEL_TEMPLATE, filepath.Join(modulePath, "model.go"))
	// route file
	RenderAndSave(model, ROUTE_TEMPLATE, filepath.Join(modulePath, "route.go"))
	RenderAndSave(model, ROUTE_IMPL_TEMPLATE, filepath.Join(implModulePath, "route.go"))
	// service file
	RenderAndSave(model, SERVICE_TEMPLATE, filepath.Join(modulePath, "service.go"))
	RenderAndSave(model, SERVICE_IMPL_TEMPLATE, filepath.Join(implModulePath, "service.go"))
	// store file
	RenderAndSave(model, STORE_TEMPLATE, filepath.Join(modulePath, "store.go"))
	RenderAndSave(model, STORE_IMPL_TEMPLATE, filepath.Join(implModulePath, "store.go"))
}

func RenderAndSave(model *ModelStruct, tmpl string, path string) {
	rendered := Render(model, tmpl, model.Name)
	if err := os.WriteFile(path, []byte(rendered), 0644); err != nil {
		panic(err)
	}
	println("generated file: " + path)
}

func Render(model *ModelStruct, tmpl string, name string) string {
	tmpl = strings.Trim(tmpl, "\n")
	engine, err := template.New(name).Parse(tmpl)

	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := engine.Execute(&buf, model); err != nil {
		panic(err)
	}
	return buf.String()
}

const MODEL_TEMPLATE = `
package {{ .Name }}

import (
	"{{ .PackageBase }}/common"
)

type {{ .Uname }} struct {
	common.BaseModel {{ range $index, $field := .Fields }} 
	{{ $field.Uname }} {{ $field.Type }} ` + "`json:\"{{ $field.Name }}\" {{ $field.ExtraTags }}`" + `{{ end }}
}

type {{ .Uname }}Querier struct {
	common.Pager {{ range $index, $field := .Fields }}
	{{ $field.Uname }} {{ $field.Type }} ` + "`json:\"{{ $field.Name }}\"`" + ` {{ end }}
}

`

const ROUTE_TEMPLATE = `
package {{ .Name }}

import (
	"github.com/gofiber/fiber/v3"
)

type I{{ .Uname }}Route interface {

	// basic crud
	Get(fiber.Ctx) error

	Save(fiber.Ctx) error

	List(fiber.Ctx) error

	Delete(fiber.Ctx) error

}

`

const ROUTE_IMPL_TEMPLATE = `
package {{ .Name }}Impl

import (
	"{{ .PackageBase }}/mods/{{ .Name }}"
	"{{ .PackageBase }}/common"
	"{{ .PackageBase }}/globals"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"context"
)

type {{ .Uname }}RouteImpl struct {
	{{ .Name }}Service {{ .Name }}.I{{ .Uname }}Service
}

func (r *{{ .Uname }}RouteImpl) Get(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("{{ .Name }} get bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	if querier.Id < 1 {
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.{{ .Name }}Service.Get(context.Background(), querier.Id)))
}
func (r *{{ .Uname }}RouteImpl) Save(c fiber.Ctx) error {
	var form {{ .Name }}.{{ .Uname }}
	if err := c.Bind().Body(&form); err != nil {
		globals.LOG.Error("{{ .Name }} save bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}	
	var result *{{ .Name }}.{{ .Uname }}
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		save, err := r.{{ .Name }}Service.Save(globals.ContextDB(context.Background(), tx), &form)
		if err != nil {
			return err
		}
		result = save
		return nil
	})
	if err != nil {
		globals.LOG.Error("exec transaction error: ", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return c.JSON(common.Ok(result))
}
func (r *{{ .Uname }}RouteImpl) List(c fiber.Ctx) error {
	var querier {{ .Name }}.{{ .Uname }}Querier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("{{ .Name }} list bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.{{ .Name }}Service.List(context.Background(), &querier)))
}
func (r *{{ .Uname }}RouteImpl) Delete(c fiber.Ctx) error {
	var querier common.BaseModel
	if err := c.Bind().URI(&querier); err != nil {
		globals.LOG.Error("{{ .Name }} delete bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	var result uint
	err := globals.DB.Transaction(func(tx *gorm.DB) error {
		id, err := r.{{ .Name }}Service.Delete(globals.ContextDB(context.Background(), tx), querier.Id)
		if err != nil {
			return err
		}
		result = id
		return nil
	})
	if err != nil {
		globals.LOG.Error("exec transaction error: ", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return c.JSON(common.Ok(result))
}
`

const SERVICE_TEMPLATE = `
package {{ .Name }}

import (
	"context"
)

type I{{ .Uname }}Service interface {

	// basic crud
	Get(context.Context, uint) (*{{ .Uname }}, error)

	Save(context.Context, *{{ .Uname }}) (*{{ .Uname }}, error)

	List(context.Context, *{{ .Uname }}Querier) ([]*{{ .Uname }}, error)

	Delete(context.Context, uint) (uint, error)

}

`

const SERVICE_IMPL_TEMPLATE = `
package {{ .Name }}Impl

import (
	"{{ .PackageBase }}/mods/{{ .Name }}"
	"{{ .PackageBase }}/store"
	"context"
)

type {{ .Uname }}Service struct {
	store store.IStore
}

func (s *{{ .Uname }}Service) Get(ctx context.Context, id uint) (*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Store := s.store.Get{{ .Uname }}Store(ctx)
	return {{ .Name }}Store.Get(id)
}

func (s *{{ .Uname }}Service) Save(ctx context.Context, form *{{ .Name }}.{{ .Uname }}) (*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Store := s.store.Get{{ .Uname }}Store(ctx)
	return {{ .Name }}Store.Save(form)
}

func (s *{{ .Uname }}Service) List(ctx context.Context, querier *{{ .Name }}.{{ .Uname }}Querier) ([]*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Store := s.store.Get{{ .Uname }}Store(ctx)
	return {{ .Name }}Store.List(querier)
}

func (s *{{ .Uname }}Service) Delete(ctx context.Context, id uint) (uint, error) {
	{{ .Name }}Store := s.store.Get{{ .Uname }}Store(ctx)
	return {{ .Name }}Store.Delete(id)
}

`

const STORE_TEMPLATE = `
package {{ .Name }}

type I{{ .Uname }}Store interface {

	// basic crud
	Get(uint) (*{{ .Uname }}, error)

	Save(*{{ .Uname }}) (*{{ .Uname }}, error)

	List(*{{ .Uname }}Querier) ([]*{{ .Uname }}, error)

	Delete(uint) (uint, error)

}

`

// TODO list implement
const STORE_IMPL_TEMPLATE = `
package {{ .Name }}Impl

import (
	"gorm.io/gorm"
	"{{ .PackageBase }}/mods/{{ .Name }}"
)

type {{ .Uname }}Store struct {
	tx *gorm.DB
}

func (s *{{ .Uname }}Store) Get(id uint) (*{{ .Name }}.{{ .Uname }}, error) {
	var model {{ .Name }}.{{ .Uname }}
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *{{ .Uname }}Store) Save(form *{{ .Name }}.{{ .Uname }}) (*{{ .Name }}.{{ .Uname }}, error) {
	if err := s.tx.Save(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *{{ .Uname }}Store) List(querier *{{ .Name }}.{{ .Uname }}Querier) ([]*{{ .Name }}.{{ .Uname }}, error) {
	var list []*{{ .Name }}.{{ .Uname }}
	if err := s.tx.Where(querier).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *{{ .Uname }}Store) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&{{ .Name }}.{{ .Uname }}{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

`
