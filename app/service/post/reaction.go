package post

import (
	"fmt"
	"log"
)

func (p *postService) DislikePost(postId, userId int) int {
	post, err := p.repository.GetPostById(int64(postId))
	if err != nil {
		log.Println(err)
		return 400
	}
	like := p.repository.GetLikeStatus(postId, userId)
	dislike := p.repository.GetDislikeStatus(postId, userId)

	if dislike == 0 && like == 0 {
		// User has not liked or disliked the post
		err = p.repository.DislikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Dislike++
	} else if dislike == 0 && like == 1 {
		// User has liked the post, changing to dislike
		err = p.repository.DeletePostLike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		err = p.repository.DislikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like--
		post.Dislike++
	} else {
		// User has already disliked the post, removing dislike
		err = p.repository.DeletePostDislike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Dislike--
	}

	err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
	if err != nil {
		log.Println(err)
		return 500
	}

	return 200
}

func (p postService) LikePost(postId, userId int) int {
	post, err := p.repository.GetPostById(int64(postId))
	if err != nil {
		log.Println(err)
		return 400
	}
	like := p.repository.GetLikeStatus(postId, userId)
	dislike := p.repository.GetDislikeStatus(postId, userId)

	if like == 0 && dislike == 0 {
		// User has not liked or disliked the post
		err = p.repository.LikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like++
	} else if like == 0 && dislike == 1 {
		// User has disliked the post, changing to like
		err = p.repository.DeletePostDislike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		err = p.repository.LikePost(postId, userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Dislike--
		post.Like++
	} else {
		// User has already liked the post, removing like
		err = p.repository.DeletePostLike(postId, userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		post.Like--
	}

	err = p.repository.UpdatePostLikeDislike(postId, int(post.Like), int(post.Dislike))
	if err != nil {
		log.Println(err)
		return 500
	}

	return 200
}

func (p postService) LikeComment(commentId, userId int) int {
	comment, err := p.repository.GetCommentByCommentID(int64(commentId))
	if err != nil {
		log.Println(err)
		return 400
	}
	like := p.repository.GetCommentLikeStatus(commentId, userId)
	dislike := p.repository.GetCommentDislikeStatus(commentId, userId)
	if like == 0 && dislike == 0 {
		err = p.repository.LikeComment(int(comment.Id), userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		comment.Like++
		err = p.repository.UpdateCommentLikeDislike(int(comment.Id), comment.Like, comment.Dislike)
		if err != nil {
			log.Println(err)
			return 500
		}
	} else if like == 0 && dislike == 1 {
		err = p.repository.DeleteCommentDislike(int(comment.Id), userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		err = p.repository.LikeComment(int(comment.Id), userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		comment.Like++
		comment.Dislike--
		err = p.repository.UpdateCommentLikeDislike(int(comment.Id), comment.Like, comment.Dislike)
		if err != nil {
			log.Println(err)
			return 500
		}
	} else {
		err = p.repository.DeleteCommentLike(int(comment.Id), userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		comment.Like--
		err = p.repository.UpdateCommentLikeDislike(int(comment.Id), comment.Like, comment.Dislike)
		if err != nil {
			log.Println(err)
			return 500
		}
	}
	return 200
}

func (p postService) DislikeComment(commentId, userId int) int {
	comment, err := p.repository.GetCommentByCommentID(int64(commentId))
	if err != nil {
		log.Println(err)
		return 400
	}
	fmt.Println("This is a comment", comment)
	like := p.repository.GetCommentLikeStatus(int(comment.Id), userId)
	dislike := p.repository.GetCommentDislikeStatus(int(comment.Id), userId)
	if like == 0 && dislike == 0 {
		err = p.repository.DislikeComment(int(comment.Id), userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		comment.Dislike++
		err = p.repository.UpdateCommentLikeDislike(int(comment.Id), comment.Like, comment.Dislike)
		if err != nil {
			log.Println(err)
			return 500
		}
	} else if dislike == 0 && like == 1 {
		err = p.repository.DeleteCommentLike(int(comment.Id), userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		err = p.repository.DislikeComment(int(comment.Id), userId, 1)
		if err != nil {
			log.Println(err)
			return 500
		}
		comment.Like--
		comment.Dislike++
		err = p.repository.UpdateCommentLikeDislike(int(comment.Id), comment.Like, comment.Dislike)
		if err != nil {
			log.Println(err)
			return 500
		}
	} else {
		err = p.repository.DeleteCommentDislike(int(comment.Id), userId)
		if err != nil {
			log.Println(err)
			return 500
		}
		comment.Dislike--
		err = p.repository.UpdateCommentLikeDislike(int(comment.Id), comment.Like, comment.Dislike)
		if err != nil {
			log.Println(err)
			return 500
		}

	}

	return 200
}
