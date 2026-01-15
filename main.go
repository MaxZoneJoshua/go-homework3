package main

import (
	"go-homework3/blog"
	"log"
)

func main() {
	// Open the database connection.
	db, err := openDB("blog.db")
	if err != nil {
		log.Fatal(err)
	}

	//db.AutoMigrate(&blog.User{}, &blog.Post{}, &blog.Comment{})
	//
	//user := blog.User{Name: "John Doe"}
	//db.Create(&user)
	//
	//db.Create(&blog.Post{Title: "First Post", UserID: user.ID})
	//db.Create(&blog.Post{Title: "Second Post", UserID: user.ID})
	//db.Create(&blog.Comment{Content: "First comment", PostID: 1})
	//db.Create(&blog.Comment{Content: "Second comment", PostID: 1})
	//db.Create(&blog.Comment{Content: "xxx comment", PostID: 2})
	//db.Create(&blog.Comment{Content: "zzzz comment", PostID: 2})

	user := blog.User{Name: "Jane Doe"}
	if err := db.Create(&user).Error; err != nil {

	}

	topPost := blog.Post{Title: "Top Post", UserID: user.ID}
	otherPost := blog.Post{Title: "Other Post", UserID: user.ID}
	if err := db.Create(&topPost).Error; err != nil {

	}
	if err := db.Create(&otherPost).Error; err != nil {
	}

	comments := []blog.Comment{
		{Content: "c1", PostID: topPost.ID},
		{Content: "c2", PostID: topPost.ID},
		{Content: "c3", PostID: topPost.ID},
		{Content: "c4", PostID: otherPost.ID},
	}
	if err := db.Create(&comments).Error; err != nil {

	}

	result, err := blog.GetPostWithMostComments(db)
	if err != nil {
	}

	if result.ID != topPost.ID {
		log.Printf("Expected post ID %d, got %d", topPost.ID, result.ID)
	}
	if result.CommentCount != 3 {
	}

	// print results
	log.Printf("Post with most comments: %+v", result)
}
