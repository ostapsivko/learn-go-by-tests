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
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Hello
world!`
		secondBody = `Title: Post 2
Description: Description 2
Tags: tag1, tag2
---
Hello world!`
	)

	t.Run("checking file contents", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewBlogPostsFromFS(fs)
		got := posts[0]

		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tag1", "tag2"},
			Body: `Hello
world!`,
		}

		assertPost(t, got, want)
	})

	t.Run("reading with failing stub", func(t *testing.T) {
		_, err := blogposts.NewBlogPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Fatal("expected an error, but got none")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
