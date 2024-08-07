package post

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"forum/app/models"
)

func (p postService) GetAllPosts() ([]models.Post, error) {
	posts, err := p.repository.GetAllPosts()
	if err != nil {
		return []models.Post{}, err
	}
	result := []models.Post{}
	for i := len(posts) - 1; i >= 0; i-- {
		result = append(result, posts[i])
	}
	return result, nil
}

func (p postService) CreatePost(post *models.Post) (int, error) {
	ok := validDataString(post.Title)
	if !ok {
		return 400, errors.New("title is invalid")
	}

	if !ok {
		return 400, errors.New("content is invalid")
	}
	ok = validCategory(post.Category)
	if !ok {
		return 400, errors.New("category is invalid")
	}
	id, err := p.repository.CreatePost(*post)
	if err != nil {
		return 500, errors.New("creating a post was failed")
	}
	categories := models.Category{
		CategoryName: post.Category,
		PostId:       id,
	}
	err = p.repository.CreateCategory(&categories)
	if err != nil {
		return 500, errors.New("creating a category was failed")
	}
	return 200, nil
}

func (p postService) GetAllCommentsAndPostsByPostId(id int64) (models.Post, int) {
	initialPost, err := p.repository.GetPostById(id)
	if err != nil {
		log.Println(err)
		return models.Post{}, 400
	}
	comments, err := p.repository.GetAllCommentByPostId(int(id))
	if err != nil {
		fmt.Println(2)
		log.Println(err)
		return models.Post{}, 500
	}

	sortedComments := []models.Comment{}

	for i := len(comments) - 1; i >= 0; i-- {
		sortedComments = append(sortedComments, comments[i])
	}
	initialPost.Comment = sortedComments
	return initialPost, 200
}

func (p postService) CreateComment(comment *models.Comment) (int, error) {
	ok := validDataString(comment.Message)
	if !ok {
		return 400, errors.New("comment message is invalid")
	}
	_, err := p.repository.GetPostById(comment.PostId)
	if err != nil {
		return 400, errors.New("post doesnt exists")
	}
	err = p.repository.CommentPost(*comment)
	if err != nil {
		log.Println(err)
		return 500, errors.New("comment post was failed")
	}
	return 200, nil
}

func validDataString(s string) bool {
	str := strings.TrimSpace(s)
	if len(str) == 0 {
		return false
	}
	//for _, v := range str {
	//	if v < rune(32) {
	//		return false
	//	}
	//}
	return true
}
