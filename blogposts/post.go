package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

type Post struct {
	Title, Description string
}

func newPost(postReader io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postReader)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readLine(titleSeparator)
	description := readLine(descriptionSeparator)

	post := Post{Title: title, Description: description}
	return post, nil
}
