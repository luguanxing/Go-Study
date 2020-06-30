package impl

type RetrieverPosterImpl struct {

}

func (*RetrieverPosterImpl) Get(url string) string {
	return "get done"
}

func (*RetrieverPosterImpl) Post(url string, form map[string]string) string {
	return "post done"
}

