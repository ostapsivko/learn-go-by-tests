package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail...")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("reading all files successfully", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := blogposts.NewBlogPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, want %d", len(posts), len(fs))
		}
	})

	t.Run("reading with failing stub", func(t *testing.T) {
		_, err := blogposts.NewBlogPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Fatal("expected an error, but got none")
		}
	})
}
