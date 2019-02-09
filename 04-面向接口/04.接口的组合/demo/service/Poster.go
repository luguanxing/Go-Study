package service

type Poster interface {
	Post(url string, form map[string] string) string
}