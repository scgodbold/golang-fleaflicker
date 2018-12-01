package nfl

import (
	"io"
	"net/http"
)

func (c *NflClient) fetchUrl(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
