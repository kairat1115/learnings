package main

import (
	blogposts "learn-go-with-tests/reading-files"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", posts)
}
