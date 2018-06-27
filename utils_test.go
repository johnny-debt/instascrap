package instascrap

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/stretchr/testify/assert"
)

// Ensures that this method returns exactly response body
func TestGetDataFromUrlSuccessful(t *testing.T) {
	defer gock.Off()

	apiUrl := "http://example.com"
	apiPath := "status"
	expectedResponse := "anything"

	gock.New(apiUrl).
		Get(apiPath).
		Reply(200).
		BodyString(expectedResponse)

	actualResponse, err := getDataFromURL(apiUrl + "/" + apiPath)

	assert.Equal(t, []byte(expectedResponse), actualResponse)
	assert.Nil(t, err)
}

func TestGetDataFromUrlNon200HttpCode(t *testing.T) {
	defer gock.Off()

	apiUrl := "http://example.com"
	apiPath := "status"

	gock.New(apiUrl).
		Get(apiPath).
		Reply(302).
		BodyString("")

	_, err := getDataFromURL(apiUrl + "/" + apiPath)

	assert.Error(t, err)
}