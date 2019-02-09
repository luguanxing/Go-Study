package retriever

type Retriever interface {
	Get(url string) string
}
