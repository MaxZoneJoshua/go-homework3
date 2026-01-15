package blog

import (
	"gorm.io/gorm"
)

func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.
		Where("user_id = ?", userID).
		Preload("Comment").
		Find(&posts).
		Error
	return posts, err
}

type PostWithCommentCount struct {
	Post
	CommentCount int64 `gorm:"column:comment_count"`
}

func GetPostWithMostComments(db *gorm.DB) (PostWithCommentCount, error) {
	var result PostWithCommentCount
	err := db.Model(&Post{}).
		Select("posts.*, count(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Scan(&result).
		Error
	return result, err
}
