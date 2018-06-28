package instascrap

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/stretchr/testify/assert"
)

// Ensures that this method returns exactly response body
func TestGetHashtagMediaSuccessful(t *testing.T) {
	defer gock.Off()
	hashtag := "something"
	maxId := ""
	apiEndpoint := "https://www.instagram.com"
	apiUri := fmt.Sprintf("explore/tags/%s", hashtag)
	params := map[string]string{"__a": "1", "max_id": maxId}

	json := ReadTestDataFile("test-01-get-medias-from-hashtag-page.json")

	gock.New(apiEndpoint).
		Get(apiUri).
		MatchParams(params).
		Reply(200).
		JSON(json)

	medias, err := GetHashtagMedia(hashtag)

	assert.NoError(t, err)
	assert.Len(t, medias, 63)
	validateMediaObjects(t, medias)
}

// Ensures that when JSON retrieving occurred this method returns error as well
func TestGetHashtagMediaJSONRetrievingError(t *testing.T) {
	defer gock.Off()
	hashtag := "something"
	maxId := ""
	apiEndpoint := "https://www.instagram.com"
	apiUri := fmt.Sprintf("explore/tags/%s", hashtag)
	params := map[string]string{"__a": "1", "max_id": maxId}

	gock.New(apiEndpoint).
		Get(apiUri).
		MatchParams(params).
		Reply(201).
		JSON("")

	medias, err := GetHashtagMedia(hashtag)

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
