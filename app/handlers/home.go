package handlers

import (
	"forum/app/models"
	"forum/pkg"
	"log"
	"net/http"
)

func (app *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pkg.ErrorHandler(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, 405)
		return
	}
	// fmt.Println(r.Context().Value(KeyUserType(keyUser)), 1111)
	// user, ok := r.Context().Value(KeyUserType(keyUser)).(models.User)
	// if !ok {
	// 	pkg.ErrorHandler(w, 401)
	// 	return
	// }
	posts, err := app.postService.GetAllPosts()
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, 500)
		return
	}
	data := models.Data{
		Posts: posts,
		// User:  user,
		Genre: "/",
	}
	pkg.RenderTemplate(w, "index.html", data)
}

func (app *App) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		pkg.ErrorHandler(w, 405)
		return
	}
	posts, err := app.postService.GetAllPosts()
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, 500)
		return
	}
	data := models.Data{
		Posts: posts,
	}
	pkg.RenderTemplate(w, "welcome.html", data)
}
