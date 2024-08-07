package handlers

import (
	"net/http"
	"strings"

	"forum/app/models"
	"forum/pkg"
)

func (app *App) FilterHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Filter called: ", r.URL.Path)
	parts := strings.Split(r.URL.Path, "/")
	category := parts[2]

	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, 405)
		return
	}
	user, ok := r.Context().Value(KeyUserType(keyUser)).(models.User)
	if !ok {
		pkg.ErrorHandler(w, 401)
		return
	}
	data, status := app.postService.GetFilterPosts(category, user)
	switch status {
	case 500:
		pkg.ErrorHandler(w, 500)
		return
	case 400:
		pkg.ErrorHandler(w, 400)
		return
	case 200:
		pkg.RenderTemplate(w, "filter.html", data)
	}
}

func (app *App) WelcomeFilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, 405)
		return
	}
	// fmt.Println("Filter called: ", r.URL.Path)
	parts := strings.Split(r.URL.Path, "/")
	category := parts[3]
	// fmt.Println("Category: ", category)
	data, status := app.postService.GetWelcomeFilterPosts(category)
	switch status {
	case 500:
		pkg.ErrorHandler(w, 500)
		return
	case 400:
		pkg.ErrorHandler(w, 400)
		return
	case 200:
		pkg.RenderTemplate(w, "welcome.html", data)
	}
}
