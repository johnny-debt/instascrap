package instascrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMediasFromHashtagPage(t *testing.T) {
	json := ReadTestDataFile("test-01-get-medias-from-hashtag-page.json")
	medias, err := getFromHashtagPage(json)
	assert.NoError(t, err, "Error is unexpected")
	// We are expecting exactly 63 medias from the given JSON response
	assert.Len(t, medias, 63, "Medias count mismatch")
}

func TestGetMediasFromHashtagPageWithWrongJson(t *testing.T) {
	_, err := getFromHashtagPage([]byte("[wrong json"))
	assert.Error(t, err)
}