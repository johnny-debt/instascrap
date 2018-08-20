package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/johnny-debt/instascrap"
	"golang.org/x/net/proxy"
)

// Provides default HTTP client (you can add proxy here or anything else)
func customHTTPClientProvider() *http.Client {
	// create http.Client with proxy
	dialer, err := proxy.SOCKS5("tcp", "111.231.88.18:1080", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial
	httpClient := &http.Client{Transport: httpTransport}
	return httpClient
}

func main() {
	is := instascrap.NewInstascrap(customHTTPClientProvider)
	fmt.Printf("Instascrap created: %v\n", is)
	getHashtagMediaExample(is)
}

// Retrieves medias list
func getHashtagMediaExample(is instascrap.Instascrap) {
	medias, err := is.GetHashtagMedia("golang")
	if err != nil {
		fmt.Printf("Medias retrieving failed, %s\n", err)
		return
	}
	fmt.Printf("Medias retrieved successfully. Count: %d", len(medias))
}
