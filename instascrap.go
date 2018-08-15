package instascrap

import (
	"errors"
	"io"
	"net/http"
)

// HTTPClientProvider is a factor
type HTTPClientProvider func() *http.Client

// Instascrap is main object which provides access to other functions
type Instascrap struct {
	httpClientProvider HTTPClientProvider
}

// NewInstascrap instantiate Instascrap structure
func NewInstascrap(httpClientProvider HTTPClientProvider) Instascrap {
	instascrap := Instascrap{}

	if httpClientProvider != nil {
		instascrap.httpClientProvider = httpClientProvider
	} else {
		instascrap.httpClientProvider = defaultHTTPClientProvider
	}

	return instascrap
}

// Provides default HTTP client
func defaultHTTPClientProvider() *http.Client {
	return &http.Client{}
}

func (instascrap *Instascrap) getDataFromURL(URL string, reader func(r io.Reader) ([]byte, error)) ([]byte, error) {
	// create a request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	// use the http client to fetch the page
	httpClient := instascrap.httpClientProvider()
	if httpClient == nil {
		return nil, errors.New("HTTP Client not  available")
	}
	resp, err := httpClient.Do(req)

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
