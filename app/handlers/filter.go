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
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	user, ok := r.Context().Value(KeyUserType(keyUser)).(models.User)
	if !ok {
		pkg.ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	data, status := app.postService.GetFilterPosts(category, user)
	switch status {
	case http.StatusInternalServerError:
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	case http.StatusBadRequest:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	case http.StatusOK:
		pkg.RenderTemplate(w, "filter.html", data)
	}
}

func (app *App) WelcomeFilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	// fmt.Println("Filter called: ", r.URL.Path)
	parts := strings.Split(r.URL.Path, "/")
	category := parts[3]
	// fmt.Println("Category: ", category)
	data, status := app.postService.GetWelcomeFilterPosts(category)
	switch status {
	case http.StatusInternalServerError:
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	case http.StatusBadRequest:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	case http.StatusOK:
		pkg.RenderTemplate(w, "welcome.html", data)
	}
}
