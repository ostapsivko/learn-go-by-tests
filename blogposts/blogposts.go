package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewBlogPostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
