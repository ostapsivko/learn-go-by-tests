package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ     *template.Template
	converter goldmark.Markdown
}

func NewPostRenderer() (*PostRenderer, error) {
	template, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	return &PostRenderer{templ: template, converter: md}, nil

}

func (r *PostRenderer) Render(writer io.Writer, post Post) error {
	var buf bytes.Buffer

	if err := r.converter.Convert([]byte(post.Body), &buf); err != nil {
		return err
	}

	post.ProcessedBody = template.HTML(buf.String())

	if err := r.templ.ExecuteTemplate(writer, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
	ProcessedBody            template.HTML
}
