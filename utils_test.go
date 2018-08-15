package instascrap

import (
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

// Ensures that this method returns exactly response body
func TestGetDataFromUrlSuccessful(t *testing.T) {
	defer gock.Off()

	apiURL := "https://example.com"
	apiPath := "status"
	expectedResponse := "anything"

	gock.New(apiURL).
		Get(apiPath).
		Reply(200).
		BodyString(expectedResponse)

	instascrap := NewInstascrap(nil)
	actualResponse, err := instascrap.getDataFromURL(apiURL+"/"+apiPath, ioutil.ReadAll)

	assert.Equal(t, []byte(expectedResponse), actualResponse)
	assert.NoError(t, err)
}

// We returns 302 redirect without location header to emulate error
func TestGetDataFromUrlError(t *testing.T) {
	defer gock.Off()

	apiURL := "http://example.com"
	apiPath := "status"

	gock.New(apiURL).
		Get(apiPath).
		Reply(302).
		BodyString("")

	instascrap := NewInstascrap(nil)
	_, err := instascrap.getDataFromURL(apiURL+"/"+apiPath, ioutil.ReadAll)

	assert.Error(t, err)
}

// We returns 201 code, which is unexpected and must produce error as well
func TestGetDataFromUrlNon200HttpCode(t *testing.T) {
	defer gock.Off()

	apiURL := "http://example.com"
	apiPath := "status"

	gock.New(apiURL).
		Get(apiPath).
		Reply(201).
		BodyString("")

	instascrap := NewInstascrap(nil)
	_, err := instascrap.getDataFromURL(apiURL+"/"+apiPath, ioutil.ReadAll)

	assert.Error(t, err)
}

// We returns 201 code, which is unexpected and must produce error as well
func TestGetDataFromUrlBodyReadError(t *testing.T) {
	defer gock.Off()

	apiURL := "http://example.com"
	apiPath := "status"

	gock.New(apiURL).
		Get(apiPath).
		Reply(200).
		BodyString("")

	instascrap := NewInstascrap(nil)
	_, err := instascrap.getDataFromURL(apiURL+"/"+apiPath, func(r io.Reader) ([]byte, error) {
		return nil, errors.New("IO Reader error occurred")
	})

	assert.Error(t, err)
}
