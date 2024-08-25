package blogposts_test

import (
	"blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewBlogPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d", len(posts), len(fs))
	}
}
