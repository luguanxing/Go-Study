package service

// 接口组合
type RetrieverPoster interface {
	Poster
	Retriever
}