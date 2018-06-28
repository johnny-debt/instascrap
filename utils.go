package instascrap

import (
	"errors"
	"io"
	"net/http"
)

func getDataFromURL(url string, reader func(r io.Reader) ([]byte, error)) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("statusCode != 200")
	}
	defer resp.Body.Close()

	body, err := reader(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
