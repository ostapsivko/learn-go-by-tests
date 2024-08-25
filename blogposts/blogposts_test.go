package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail...")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("checking file contents", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, _ := blogposts.NewBlogPostsFromFS(fs)

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("reading with failing stub", func(t *testing.T) {
		_, err := blogposts.NewBlogPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Fatal("expected an error, but got none")
		}
	})
}
