package tmdbgae

import (
	"fmt"
	"golang.org/x/net/context"
)

// Keyword struct
type Keyword struct {
	ID   int
	Name string
}

// GetKeywordInfo gets the basic information for a specific keyword id
// http://docs.themoviedb.apiary.io/#reference/keywords/keywordid/get
func (tmdb *TMDb) GetKeywordInfo(ctx context.Context, id int) (*Keyword, error) {
	var keywordInfo Keyword
	uri := fmt.Sprintf("%s/keyword/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(ctx, uri, &keywordInfo)
	return result.(*Keyword), err
}

// GetKeywordMovies gets the list of movies for a particular keyword by id
// http://docs.themoviedb.apiary.io/#reference/keywords/keywordidmovies/get
func (tmdb *TMDb) GetKeywordMovies(ctx context.Context, id int, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"language": {},
		"page":     {}}
	var movies MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/keyword/%v/movies?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(ctx, uri, &movies)
	return result.(*MoviePagedResults), err
}
