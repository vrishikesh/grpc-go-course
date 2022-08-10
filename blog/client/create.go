package main

import (
	"context"
	pb "grpc-go-course/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating the blog")

	blog := &pb.Blog{
		AuthorId: "Author ID",
		Title:    "Title",
		Content:  "Content",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while creating the blog: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", createBlogRes.Id)
	return createBlogRes.Id
}
