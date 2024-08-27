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
	_, err := fmt.Fprintf(writer, "<h1>%s</h1>\n", post.Title)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "<p>%s</p>\n", post.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(writer, "Tags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range post.Tags {
		_, err = fmt.Fprintf(writer, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(writer, "</ul>")
	if err != nil {
		return err
	}

	return nil
}
