package blogrenderer_test

import (
	"blogrenderer"
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello World!",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"tdd", "go"},
		}
	)

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello World!</h1>
<p>This is a description</p>
Tags: <ul><li>tdd</li><li>go</li></ul>`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
