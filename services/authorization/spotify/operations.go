package spotify

type Operations struct {
	url string
}

func NewOperations(url string) Operations {
	return Operations{
		url: url,
	}
}
