package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title, Description string
	Tags               []string
}

func newPost(postReader io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postReader)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readLine(titleSeparator)
	description := readLine(descriptionSeparator)
	tagsLine := readLine(tagsSeparator)

	tags := strings.Split(tagsLine, ", ")

	post := Post{Title: title, Description: description, Tags: tags}
	return post, nil
}
