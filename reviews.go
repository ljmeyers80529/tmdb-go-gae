package tmdbgae

import (
	"fmt"
	"golang.org/x/net/context"
)

// Review struct
type Review struct {
	ID         string
	Author     string
	Content    string
	Iso639_1   string `json:"iso_639_1"`
	MediaID    int    `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
	URL        string
}

// GetReviewInfo gets the full details of a review by ID
// http://docs.themoviedb.apiary.io/#reference/reviews/reviewid/get
func (tmdb *TMDb) GetReviewInfo(ctx context.Context, id string) (*Review, error) {
	var reviewInfo Review
	uri := fmt.Sprintf("%s/review/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(ctx, uri, &reviewInfo)
	return result.(*Review), err
}
