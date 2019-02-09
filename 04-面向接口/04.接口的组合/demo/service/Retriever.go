package service

type Retriever interface {
	Get(url string) string
}
