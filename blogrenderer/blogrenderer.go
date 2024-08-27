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

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(writer io.Writer, post Post) error {
	template, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	if err := template.Execute(writer, post); err != nil {
		return err
	}

	return nil
}
