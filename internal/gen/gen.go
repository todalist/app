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



// a naive code generator.
// 
// TODO dryrun support
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
		// TODO ask user to confirm
		println("previous module will be deleted: " + model.Name)
		if err := os.RemoveAll(modulePath); err != nil {
			panic(err)
		}
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
	// repo file
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
	engine := template.New(name)
	engine.Funcs(template.FuncMap{
		"isPtrType": func(t string) bool {
			return strings.HasPrefix(t, "*")
		},
	})
	engine, err := engine.Parse(tmpl)

	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := engine.Execute(&buf, model); err != nil {
		panic(err)
	}
	return buf.String()
}

