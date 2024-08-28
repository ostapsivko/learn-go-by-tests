package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	template, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: template}, nil

}

func (r *PostRenderer) Render(writer io.Writer, post Post) error {

	if err := r.templ.ExecuteTemplate(writer, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
}
