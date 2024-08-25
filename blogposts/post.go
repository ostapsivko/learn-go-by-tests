package blogposts

import "io"

type Post struct {
	Title string
}

func newPost(postReader io.Reader) (Post, error) {
	postData, err := io.ReadAll(postReader)

	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData[7:])}
	return post, nil
}
