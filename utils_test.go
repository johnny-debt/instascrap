package instascrap

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io"
	"io/ioutil"
	"testing"
)

// Ensures that this method returns exactly response body
func TestGetDataFromUrlSuccessful(t *testing.T) {
	defer gock.Off()

	apiUrl := "https://example.com"
	apiPath := "status"
	expectedResponse := "anything"

	gock.New(apiUrl).
		Get(apiPath).
		Reply(200).
		BodyString(expectedResponse)

	actualResponse, err := getDataFromURL(apiUrl+"/"+apiPath, ioutil.ReadAll)

	assert.Equal(t, []byte(expectedResponse), actualResponse)
	assert.NoError(t, err)
}

// We returns 302 redirect without location header to emulate error
func TestGetDataFromUrlError(t *testing.T) {
	defer gock.Off()

	apiUrl := "http://example.com"
	apiPath := "status"

	gock.New(apiUrl).
		Get(apiPath).
		Reply(302).
		BodyString("")

	_, err := getDataFromURL(apiUrl+"/"+apiPath, ioutil.ReadAll)

	assert.Error(t, err)
}

// We returns 201 code, which is unexpected and must produce error as well
func TestGetDataFromUrlNon200HttpCode(t *testing.T) {
	defer gock.Off()

	apiUrl := "http://example.com"
	apiPath := "status"

	gock.New(apiUrl).
		Get(apiPath).
		Reply(201).
		BodyString("")

	_, err := getDataFromURL(apiUrl+"/"+apiPath, ioutil.ReadAll)

	assert.Error(t, err)
}

// We returns 201 code, which is unexpected and must produce error as well
func TestGetDataFromUrlBodyReadError(t *testing.T) {
	defer gock.Off()

	apiUrl := "http://example.com"
	apiPath := "status"

	gock.New(apiUrl).
		Get(apiPath).
		Reply(200).
		BodyString("")

	_, err := getDataFromURL(apiUrl+"/"+apiPath, func(r io.Reader) ([]byte, error) {
		return nil, errors.New("IO Reader error occurred")
	})

	assert.Error(t, err)
}
