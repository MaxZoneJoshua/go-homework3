package blog

import "gorm.io/gorm"

type User struct {
	ID uint `gorm:"primary_key"`

	Name string

	Post []Post

	PostCount int `gorm:"default:0"`
}

type Post struct {
	ID uint `gorm:"primary_key"`

	Title string

	// 默认外键 type + primaryKey 命名
	UserID uint

	Comment []Comment

	CommentStatus string `gorm:"default:'无评论'"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	if p.UserID == 0 {
		return nil
	}

	tx.Model(&Post{}).
		Where("id = ?", p.ID).
		Update("comment_status", "有评论")

	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).
		Error
}

type Comment struct {
	ID uint `gorm:"primary_key"`

	Content string

	PostID uint
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	if c.PostID == 0 {
		return nil
	}

	var count int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	return tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_status", "无评论").
		Error
}
