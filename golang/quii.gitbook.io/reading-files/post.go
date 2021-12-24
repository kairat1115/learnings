package readingfiles

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(r io.Reader) (Post, error) {
	s := bufio.NewScanner(r)

	readMetaLine := func(tagName string) string {
		s.Scan()
		return strings.TrimPrefix(s.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(s),
	}, nil
}

func readBody(s *bufio.Scanner) string {
	s.Scan() // ignore ---
	buf := bytes.Buffer{}
	for s.Scan() {
		fmt.Fprintln(&buf, s.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
