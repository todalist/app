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

	if f, err := os.Stat(modulePath); err == nil && f.IsDir() {
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
	// Repo file
	RenderAndSave(model, REPO_TEMPLATE, filepath.Join(modulePath, "repo.go"))
	RenderAndSave(model, REPO_IMPL_TEMPLATE, filepath.Join(implModulePath, "repo.go"))
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

	First(fiber.Ctx) error

	Save(fiber.Ctx) error

	List(fiber.Ctx) error

	Delete(fiber.Ctx) error

	Register(fiber.Router)

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

func (r *{{ .Uname }}RouteImpl) First(c fiber.Ctx) error {
	var querier {{ .Name }}.{{ .Uname }}Querier
	if err := c.Bind().Body(&querier); err != nil {
		globals.LOG.Error("{{ .Name }} first bind error", zap.String("error", err.Error()))
		return fiber.ErrBadRequest
	}
	return c.JSON(common.Or(r.{{ .Name }}Service.First(context.Background(), &querier)))
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

func (r *{{ .Uname }}RouteImpl) Register(root fiber.Router) {
	router := root.Group("/{{ .Name }}")
	router.Get("/:id", r.Get)
	router.Post("/save", r.Save)
	router.Post("/list", r.List)
	router.Delete("/:id", r.Delete)
}

func New{{ .Uname }}Route({{ .Name }}Service {{ .Name }}.I{{ .Uname }}Service) {{ .Name }}.I{{ .Uname }}Route {
	return &{{ .Uname }}RouteImpl{
		{{ .Name }}Service: {{ .Name }}Service,
	}
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

	First(context.Context, *{{ .Uname }}Querier) (*{{ .Uname }}, error)

	Save(context.Context, *{{ .Uname }}) (*{{ .Uname }}, error)

	List(context.Context, *{{ .Uname }}Querier) ([]*{{ .Uname }}, error)

	Delete(context.Context, uint) (uint, error)

}

`

const SERVICE_IMPL_TEMPLATE = `
package {{ .Name }}Impl

import (
	"{{ .PackageBase }}/mods/{{ .Name }}"
	"{{ .PackageBase }}/repo"
	"context"
)

type {{ .Uname }}Service struct {
	repo repo.IRepo
}

func (s *{{ .Uname }}Service) Get(ctx context.Context, id uint) (*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Repo := s.repo.Get{{ .Uname }}Repo(ctx)
	return {{ .Name }}Repo.Get(id)
}

func (s *{{ .Uname }}Service) First(ctx context.Context, querier *{{ .Name }}.{{ .Uname }}Querier) (*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Repo := s.repo.Get{{ .Uname }}Repo(ctx)
	return {{ .Name }}Repo.First(querier)
}

func (s *{{ .Uname }}Service) Save(ctx context.Context, form *{{ .Name }}.{{ .Uname }}) (*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Repo := s.repo.Get{{ .Uname }}Repo(ctx)
	return {{ .Name }}Repo.Save(form)
}

func (s *{{ .Uname }}Service) List(ctx context.Context, querier *{{ .Name }}.{{ .Uname }}Querier) ([]*{{ .Name }}.{{ .Uname }}, error) {
	{{ .Name }}Repo := s.repo.Get{{ .Uname }}Repo(ctx)
	return {{ .Name }}Repo.List(querier)
}

func (s *{{ .Uname }}Service) Delete(ctx context.Context, id uint) (uint, error) {
	{{ .Name }}Repo := s.repo.Get{{ .Uname }}Repo(ctx)
	return {{ .Name }}Repo.Delete(id)
}

func New{{ .Uname }}Service(repo repo.IRepo) *{{ .Uname }}Service {
	return &{{ .Uname }}Service{
		repo: repo,
	}
}

`

const REPO_TEMPLATE = `
package {{ .Name }}

type I{{ .Uname }}Repo interface {

	// basic crud
	Get(uint) (*{{ .Uname }}, error)

	First(*{{ .Uname }}Querier) (*{{ .Uname }}, error)

	Save(*{{ .Uname }}) (*{{ .Uname }}, error)

	List(*{{ .Uname }}Querier) ([]*{{ .Uname }}, error)

	Delete(uint) (uint, error)

}

`

// TODO list implement
const REPO_IMPL_TEMPLATE = `
package {{ .Name }}Impl

import (
	"gorm.io/gorm"
	"{{ .PackageBase }}/mods/{{ .Name }}"
	"{{ .PackageBase }}/common"
)

type {{ .Uname }}Repo struct {
	tx *gorm.DB
}

func (s *{{ .Uname }}Repo) Get(id uint) (*{{ .Name }}.{{ .Uname }}, error) {
	var model {{ .Name }}.{{ .Uname }}
	if err := s.tx.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *{{ .Uname }}Repo) First(querier *{{ .Name }}.{{ .Uname }}Querier) (*{{ .Name }}.{{ .Uname }}, error) {
	list, err := s.List(querier)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (s *{{ .Uname }}Repo) Save(form *{{ .Name }}.{{ .Uname }}) (*{{ .Name }}.{{ .Uname }}, error) {
	if form.Id == 0 {
		if err := s.tx.Create(form).Error; err != nil {
			return nil, err
		}
		return form, nil
	}
	if err := s.
		tx.
		Model(form).
		Where("id = ?", form.Id).
		Where("updated_at <=", form.UpdatedAt).{{ if gt (len .UpdateOmits) 0 }}Omit({{ range  $index, $omit := .UpdateOmits }}"{{ $omit }}", {{ end }}).{{ end }}
		Updates(form).Error; err != nil {
		return nil, err
	}
	return form, nil
}

func (s *{{ .Uname }}Repo) List(querier *{{ .Name }}.{{ .Uname }}Querier) ([]*{{ .Name }}.{{ .Uname }}, error) {
	var list []*{{ .Name }}.{{ .Uname }}
	sql := s.tx.Where(querier)
	sql = common.Paginate(sql, &querier.Pager)
	if err := sql.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *{{ .Uname }}Repo) Delete(id uint) (uint, error) {
	if err := s.tx.Where("id = ?", id).Delete(&{{ .Name }}.{{ .Uname }}{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func New{{ .Uname }}Repo(tx *gorm.DB) {{ .Name }}.I{{ .Uname }}Repo {
	return &{{ .Uname }}Repo{
		tx: tx,
	}
}

`
