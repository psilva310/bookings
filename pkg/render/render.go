package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"pkg/config"
	"pkg/models"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td // for now, this is doing nothing
}

var functions = template.FuncMap{}

var app *config.AppConfig

//get the value from the passed address and insert into a global var
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	//get a specific template
	template, exist := templateCache[tmpl]
	// if the requested template does not exist
	if !exist {
		log.Fatal("Could not get template from templateCache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = template.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//return the path from matched files
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		//trim the name from path
		name := filepath.Base(page)

		//create a new template with name, functions, and parses the path
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		//return the path from matched files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			// join *.layout.tmpl with *.pages.tmpl
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
