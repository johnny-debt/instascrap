package instascrap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

// Ensures that this method returns exactly response body
func TestGetHashtagMediaSuccessful(t *testing.T) {
	defer gock.Off()
	hashtag := "something"
	maxID := ""
	apiEndpoint := "https://www.instagram.com"
	apiURI := fmt.Sprintf("explore/tags/%s", hashtag)
	params := map[string]string{"__a": "1", "max_id": maxID}

	json := ReadTestDataFile("test-01-get-medias-from-hashtag-page.json")

	gock.New(apiEndpoint).
		Get(apiURI).
		MatchParams(params).
		Reply(200).
		JSON(json)

	instascrap := NewInstascrap(nil)
	medias, err := instascrap.GetHashtagMedia(hashtag)

	assert.NoError(t, err)
	assert.Len(t, medias, 63)
	validateMediaObjects(t, medias)
}

// Ensures that when JSON retrieving occurred this method returns error as well
func TestGetHashtagMediaJSONRetrievingError(t *testing.T) {
	defer gock.Off()
	hashtag := "something"
	maxID := ""
	apiEndpoint := "https://www.instagram.com"
	apiURI := fmt.Sprintf("explore/tags/%s", hashtag)
	params := map[string]string{"__a": "1", "max_id": maxID}

	gock.New(apiEndpoint).
		Get(apiURI).
		MatchParams(params).
		Reply(201).
		JSON("")

	instascrap := NewInstascrap(nil)
	medias, err := instascrap.GetHashtagMedia(hashtag)

	assert.Error(t, err)
	assert.Nil(t, medias)
}

func validateMediaObjects(t *testing.T, medias []Media) {
	for _, media := range medias {
		validateMediaObject(t, media)
	}
}

func validateMediaObject(t *testing.T, media Media) {
	assert.NotEmpty(t, media.ID)
	assert.NotEmpty(t, media.Shortcode)
}
