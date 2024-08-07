package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/app/models"
	"forum/pkg"
)

func (app *App) CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		parts := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, 404)
			return
		}
		initialPost, status := app.postService.GetAllCommentsAndPostsByPostId(int64(id))
		switch status {
		case 405:
			pkg.ErrorHandler(w, 500)
			return
		case 400:
			pkg.ErrorHandler(w, 400)
			return
		}

		data := models.Data{
			Comment:     initialPost.Comment,
			InitialPost: initialPost,
		}

		pkg.RenderTemplate(w, "commentview.html", data)
	case http.MethodPost:
		parts := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(parts[3])
		fmt.Println(id)
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, 404)
			return
		}
		message := r.FormValue("comment")
		path := "/post/comment/" + parts[3]
		user, ok := r.Context().Value(KeyUserType(keyUser)).(models.User)
		if !ok {
			pkg.ErrorHandler(w, 401) // 401-status unauthorized
			return
		}
		comment := models.Comment{
			PostId:   int64(id),
			UserId:   user.ID,
			Username: user.Username,
			Message:  message,
			Born:     time.Now().Format(time.RFC822),
		}

		status, err := app.postService.CreateComment(&comment)
		if err != nil {
			log.Println(err)
		}
		switch status {
		case 500:
			pkg.ErrorHandler(w, 500)
			return
		case 400:
			pkg.ErrorHandler(w, 400)
			return
		case 200:
			http.Redirect(w, r, path, 302)
		}
	default:
		pkg.ErrorHandler(w, 405)
	}
}

func (app *App) WelcomeCommentHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("welcome to welcome")
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, 405)
		return
	}
	parts := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, 400)
		return
	}
	initialPost, status := app.postService.GetAllCommentsAndPostsByPostId(int64(id))
	switch status {
	case 500:
		pkg.ErrorHandler(w, 500)
		return
	case 400:
		pkg.ErrorHandler(w, 400)
		return
	}
	data := models.Data{
		Comment:     initialPost.Comment,
		InitialPost: initialPost,
	}
	pkg.RenderTemplate(w, "commentunauth.html", data)
}
