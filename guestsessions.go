package tmdbgae

import (
	"fmt"
	"golang.org/x/net/context"
)

// GetGuestSessionRatedMovies gets the list of rated movies for a specific guest session id
// http://docs.themoviedb.apiary.io/#reference/guest-sessions/guestsessionguestsessionidratedmovies/get
func (tmdb *TMDb) GetGuestSessionRatedMovies(ctx context.Context, sessionID string, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":       {},
		"sort_by":    {},
		"sort_order": {},
		"language":   {}}
	var favorites MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/guest_session/%v/rated_movies?api_key=%s%s", baseURL, sessionID, tmdb.apiKey, optionsString)
	result, err := getTmdb(ctx, uri, &favorites)
	return result.(*MoviePagedResults), err
}
