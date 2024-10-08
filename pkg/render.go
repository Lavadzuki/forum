package pkg

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"forum/app/models"
)

var (
	templateCache = make(map[string]*template.Template)
	once          sync.Once
)

func RenderTemplate(w http.ResponseWriter, template string, data models.Data) {
	once.Do(func() {
		err := createTemplate()
		if err != nil {
			log.Println(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	})

	t, ok := templateCache[template]
	if !ok {

		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func createTemplate() error {
	pages, err := filepath.Glob("./templates/html/*.html")
	if err != nil {
		return err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return err
		}
		templateCache[name] = ts
	}
	return nil
}
