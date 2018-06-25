package instascrap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMediasFromHashtagPage(t *testing.T)  {
	json := ReadTestDataFile("test-01-get-medias-from-hashtag-page.json")
	medias, err := getFromHashtagPage(json)
	assert.Nil(t, err, "Error is unexpected")
	// We are expecting exactly 63 medias from the given JSON response
	assert.Len(t, medias, 63, "Medias count mismatch")
}