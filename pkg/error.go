package pkg

import (
	"html/template"
	"log"
	"net/http"
)

type errorData struct {
	StatusText string
	StatusCode int
}

func ErrorHandler(w http.ResponseWriter, status int) {
	data := errorData{
		StatusText: http.StatusText(status),
		StatusCode: status,
	}
	tmpl, err := template.ParseFiles("./templates/html/error.html")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	w.WriteHeader(status)
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}

}
