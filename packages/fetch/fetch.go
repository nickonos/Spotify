package fetch

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type FetchResponse[TBody any] struct {
	StatusCode int
	bytes      []byte
}

func (f FetchResponse[TBody]) ReadBody() (*TBody, error) {
	var body TBody
	err := json.Unmarshal(f.bytes, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

var client *http.Client

func init() {
	client = &http.Client{}
}

func Fetch[TBody any](method string, url string, body io.Reader) (*FetchResponse[TBody], error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &FetchResponse[TBody]{
		StatusCode: req.Response.StatusCode,
		bytes:      bytes,
	}, nil
}

func FetchContext[TBody any](ctx context.Context, method string, url string, body io.Reader) (*FetchResponse[TBody], error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &FetchResponse[TBody]{
		StatusCode: req.Response.StatusCode,
		bytes:      bytes,
	}, nil
}
