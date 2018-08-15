package instascrap

import (
	"fmt"
	"io/ioutil"
)

const hashtagMediasURL = "https://www.instagram.com/explore/tags/%s/?__a=1&max_id=%s"

// A Media describes an Instagram media info.
type Media struct {
	ID            string
	Owner         Account
	Shortcode     string
	Date          uint64
	Caption       string
	CommentsCount uint32
	LikesCount    uint32
	IsAdvertising bool
}

// GetHashtagMedia returns latest medias from the hashtag page
func (instascrap *Instascrap) GetHashtagMedia(tag string) ([]Media, error) {
	var medias []Media
	url := fmt.Sprintf(hashtagMediasURL, tag, "")
	jsonBody, err := instascrap.getDataFromURL(url, ioutil.ReadAll)
	if err != nil {
		return nil, err
	}

	medias, err = getFromHashtagPage(jsonBody)
	return medias, nil
}
