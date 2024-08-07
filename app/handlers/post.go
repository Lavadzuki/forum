package handlers

import (
	"log"
	"net/http"
	"time"

	"forum/app/models"
	"forum/pkg"
)

func (app *App) PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pkg.RenderTemplate(w, "createpost.html", models.Data{})
		return
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			return
		}
		title := r.FormValue("title")
		message := r.FormValue("message")
		genre := r.Form["category"]
		user, ok := r.Context().Value(KeyUserType(keyUser)).(models.User)
		if !ok {
			pkg.ErrorHandler(w, 401)
			return
		}
		post := models.Post{
			Title:       title,
			Content:     message,
			Category:    models.Stringslice(genre),
			Author:      user,
			CreatedTime: time.Now().Format(time.RFC822),
		}
		status, err := app.postService.CreatePost(&post)
		if err != nil {
			log.Println(err)
			switch status {
			case 500:
				pkg.ErrorHandler(w, 500)
				return
			case 400:
				pkg.ErrorHandler(w, 400)
				return
			}
		}
		http.Redirect(w, r, "/", 302)

	default:
		pkg.ErrorHandler(w, 405)
		return
	}
}
