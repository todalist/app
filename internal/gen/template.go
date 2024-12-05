package main

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
	common.Pager 
	Id *uint ` + "`json:\"id\"`" + ` {{ range $index, $field := .Fields }}
	{{ $field.Uname }} {{ if isPtrType $field.Type }}{{else}}*{{ end }}{{ $field.Type }} ` + "`json:\"{{ $field.Name }}\"`" + ` {{ end }}
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
	router.Post("/first", r.First)
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
	return s.First(&{{ .Name }}.{{ .Uname }}Querier{Id: &id})
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

