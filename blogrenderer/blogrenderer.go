package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(writer io.Writer, post Post) error {
	fmt.Fprintf(writer, "<h1>%s</h1>", post.Title)
	return nil
}
