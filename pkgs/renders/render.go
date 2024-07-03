package renders

import (
	"bytes"
	"github.com/sungatishe/pkgs/config"
	"github.com/sungatishe/pkgs/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCheck
	} else {
		tc, _ = CreateCacheTemplate()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("err")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateCacheTemplate() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("../templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../templates/*layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil

}
