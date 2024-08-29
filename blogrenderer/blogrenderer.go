package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/yuin/goldmark"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostViewModel struct {
	Post
	HTMLBody template.HTML
}

type PostRenderer struct {
	templ     *template.Template
	converter goldmark.Markdown
}

func NewPostRenderer() (*PostRenderer, error) {
	template, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	md := goldmark.New()

	return &PostRenderer{templ: template, converter: md}, nil

}

func (r *PostRenderer) Render(writer io.Writer, post Post) error {
	var buf bytes.Buffer

	if err := r.converter.Convert([]byte(post.Body), &buf); err != nil {
		return err
	}

	pvm := NewViewModel(post)

	pvm.HTMLBody = template.HTML(buf.String())

	if err := r.templ.ExecuteTemplate(writer, "blog.gohtml", pvm); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(writer io.Writer, posts []Post) error {
	var viewModels []PostViewModel

	for _, post := range posts {
		viewModels = append(viewModels, NewViewModel(post))
	}

	if err := r.templ.ExecuteTemplate(writer, "index.gohtml", viewModels); err != nil {
		return err
	}

	return nil
}

func NewViewModel(post Post) PostViewModel {
	return PostViewModel{Post: post}
}

func (pvm *PostViewModel) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(pvm.Title, " ", "-", -1))
}
