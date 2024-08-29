package blogrenderer_test

import (
	"blogrenderer"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

const (
	body = "BUILD\n" +
		"- \n" +
		"Go to main directory. Run `dotnet run` in terminal.\n" +
		"\n" +
		"Hi! I'm your first Markdown file in **StackEdit**. \n" +
		"If you want to learn about StackEdit, you can read me.\n" +
		"If you want to play with Markdown, you can edit me.\n" +
		"Once you have finished with me, you can create new files by opening the **file explorer** on the left corner of the navigation bar."
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello World!",
			Body:        body,
			Description: "This is a description",
			Tags:        []string{"tdd", "go"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRenderer(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello World!",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"tdd", "go"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
