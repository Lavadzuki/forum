package handlers

import (
	"fmt"
	"forum/app/models"
	"forum/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (app *App) ReactionHandler(w http.ResponseWriter, r *http.Request) {
	path := ""
	ID := 0
	commentID := 0
	isMainPage := r.FormValue("isMainPage")
	category := r.FormValue("FILTER")
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, 405)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if parts[2] == "like" || parts[2] == "dislike" {
		id, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, 500)
			return
		}
		ID = id
		path = "/" + parts[1] + "/" + parts[2]
	} else if parts[3] == "like" || parts[3] == "dislike" {
		id, err := strconv.Atoi(parts[4])
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, 500)
			return
		}
		ID = id
		path = "/" + parts[1] + "/" + parts[2] + "/" + parts[3]

		commentiD, err := strconv.Atoi(parts[5])
		if err != nil {
			log.Println(err)
			pkg.ErrorHandler(w, 500)
			return
		}
		commentID = commentiD
	}
	user, ok := r.Context().Value(KeyUserType(keyUser)).(models.User)
	if !ok {
		pkg.ErrorHandler(w, 401)
		return
	}

	fmt.Println("Switch", path)

	switch path {
	case "/post/like":
		status := app.postService.LikePost(ID, int(user.ID))
		switch status {
		case 500:
			pkg.ErrorHandler(w, 500)
			return
		case 400:
			pkg.ErrorHandler(w, 400)
			return
		case 200:

			switch category {
			case "liked-post":
				http.Redirect(w, r, "/filter/liked-post/", 302)
			case "created-post":
				http.Redirect(w, r, "/filter/created-post/", 302)
			case "romance":
				http.Redirect(w, r, "/filter/romance/", 302)
			case "adventure":
				http.Redirect(w, r, "/filter/adventure/", 302)
			case "comedy":
				http.Redirect(w, r, "/filter/comedy/", 302)
			case "drama":
				http.Redirect(w, r, "/filter/drama/", 302)
			case "fantasy":
				http.Redirect(w, r, "/filter/fantasy/", 302)
			}

			if isMainPage == "true" {
				http.Redirect(w, r, "/", 302)
			} else {
				http.Redirect(w, r, "/post/comment/"+strconv.Itoa(ID), 302)
			}
		}
	case "/post/dislike":
		status := app.postService.DislikePost(ID, int(user.ID))
		switch status {
		case 500:
			pkg.ErrorHandler(w, 500)
			return
		case 400:
			pkg.ErrorHandler(w, 400)
			return
		case 200:
			switch category {
			case "liked-post":
				http.Redirect(w, r, "/filter/liked-post/", 302)
			case "created-post":
				http.Redirect(w, r, "/filter/created-post/", 302)
			case "romance":
				http.Redirect(w, r, "/filter/romance/", 302)
			case "adventure":
				http.Redirect(w, r, "/filter/adventure/", 302)
			case "comedy":
				http.Redirect(w, r, "/filter/comedy/", 302)
			case "drama":
				http.Redirect(w, r, "/filter/drama/", 302)
			case "fantasy":
				http.Redirect(w, r, "/filter/fantasy/", 302)
			}

			if isMainPage == "true" {
				http.Redirect(w, r, "/", 302)
			} else {
				http.Redirect(w, r, "/post/comment/"+strconv.Itoa(ID), 302)
			}
		}
	case "/post/comment/like":
		status := app.postService.LikeComment(commentID, int(user.ID))
		switch status {
		case 500:
			pkg.ErrorHandler(w, 500)
			return
		case 400:
			pkg.ErrorHandler(w, 400)
			return
		case 200:
			http.Redirect(w, r, "/post/comment/"+strconv.Itoa(ID), 302)
		}
	case "/post/comment/dislike":
		status := app.postService.DislikeComment(commentID, int(user.ID))
		switch status {
		case 500:
			pkg.ErrorHandler(w, 500)
			return
		case 400:
			pkg.ErrorHandler(w, 400)
			return
		case 200:
			http.Redirect(w, r, "/post/comment/"+strconv.Itoa(ID), 302)
		}

	default:
		pkg.ErrorHandler(w, 404)
	}
}
