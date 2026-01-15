package test

import (
	"go-homework3/blog"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}

	if err := db.AutoMigrate(&blog.User{}, &blog.Post{}, &blog.Comment{}); err != nil {
		t.Fatalf("auto migrate: %v", err)
	}

	return db
}

func TestGetUserPostsWithComments(t *testing.T) {
	db := setupDB(t)

	user := blog.User{Name: "John Doe"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}

	postOne := blog.Post{Title: "First Post", UserID: user.ID}
	postTwo := blog.Post{Title: "Second Post", UserID: user.ID}
	if err := db.Create(&postOne).Error; err != nil {
		t.Fatalf("create post one: %v", err)
	}
	if err := db.Create(&postTwo).Error; err != nil {
		t.Fatalf("create post two: %v", err)
	}

	comments := []blog.Comment{
		{Content: "Nice", PostID: postOne.ID},
		{Content: "Great", PostID: postOne.ID},
		{Content: "Thanks", PostID: postTwo.ID},
	}
	if err := db.Create(&comments).Error; err != nil {
		t.Fatalf("create comments: %v", err)
	}

	posts, err := blog.GetUserPostsWithComments(db, user.ID)
	if err != nil {
		t.Fatalf("GetUserPostsWithComments: %v", err)
	}

	if len(posts) != 2 {
		t.Fatalf("expected 2 posts, got %d", len(posts))
	}

	commentCounts := map[uint]int{}
	for _, post := range posts {
		commentCounts[post.ID] = len(post.Comment)
	}

	if commentCounts[postOne.ID] != 2 {
		t.Fatalf("postOne expected 2 comments, got %d", commentCounts[postOne.ID])
	}
	if commentCounts[postTwo.ID] != 1 {
		t.Fatalf("postTwo expected 1 comment, got %d", commentCounts[postTwo.ID])
	}
}

func TestGetPostWithMostComments(t *testing.T) {
	db := setupDB(t)

	user := blog.User{Name: "Jane Doe"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}

	topPost := blog.Post{Title: "Top Post", UserID: user.ID}
	otherPost := blog.Post{Title: "Other Post", UserID: user.ID}
	if err := db.Create(&topPost).Error; err != nil {
		t.Fatalf("create top post: %v", err)
	}
	if err := db.Create(&otherPost).Error; err != nil {
		t.Fatalf("create other post: %v", err)
	}

	comments := []blog.Comment{
		{Content: "c1", PostID: topPost.ID},
		{Content: "c2", PostID: topPost.ID},
		{Content: "c3", PostID: topPost.ID},
		{Content: "c4", PostID: otherPost.ID},
	}
	if err := db.Create(&comments).Error; err != nil {
		t.Fatalf("create comments: %v", err)
	}

	result, err := blog.GetPostWithMostComments(db)
	if err != nil {
		t.Fatalf("GetPostWithMostComments: %v", err)
	}

	if result.ID != topPost.ID {
		t.Fatalf("expected top post id %d, got %d", topPost.ID, result.ID)
	}
	if result.CommentCount != 3 {
		t.Fatalf("expected 3 comments, got %d", result.CommentCount)
	}

	// print results
	t.Logf("Post with most comments: %+v", result)
}
