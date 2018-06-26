package instascrap

import "encoding/json"

type account struct {
	ID string `json:"id"`
}

// Retrieves medias from JSON returned by an hashtag page
func getFromHashtagPage(data []byte) ([]Media, error) {
	var mediasJSON struct {
		Graphql struct {
			Hashtag struct {
				EdgeHashtagToMedia struct {
					Edges []struct {
						Node struct {
							ID                 string  `json:"id"`
							Shortcode          string  `json:"shortcode"`
							Owner              account `json:"owner"`
							EdgeMediaToCaption struct {
								Edges []struct {
									Node struct {
										Text string `json:"text"`
									} `json:"node"`
								} `json:"edges"`
							} `json:"edge_media_to_caption"`
							EdgeMediaToComment struct {
								Count uint32 `json:"count"`
							} `json:"edge_media_to_comment"`
							EdgeLikedBy struct {
								Count uint32 `json:"count"`
							} `json:"edge_liked_by"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"edge_hashtag_to_media"`
			} `json:"hashtag"`
		} `json:"graphql"`
	}

	err := json.Unmarshal(data, &mediasJSON)
	if err != nil {
		return []Media{}, err
	}

	var medias []Media

	for _, mediaJSON := range mediasJSON.Graphql.Hashtag.EdgeHashtagToMedia.Edges {
		media := Media{}
		media.ID = mediaJSON.Node.ID
		media.Shortcode = mediaJSON.Node.Shortcode
		media.Owner = Account{ID: mediaJSON.Node.Owner.ID}
		for _, captionEdge := range mediaJSON.Node.EdgeMediaToCaption.Edges {
			media.Caption += captionEdge.Node.Text
		}
		media.CommentsCount = mediaJSON.Node.EdgeMediaToComment.Count

		medias = append(medias, media)
	}

	return medias, nil
}
