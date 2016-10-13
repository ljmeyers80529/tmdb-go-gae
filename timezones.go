package tmdbgae

import (
	"fmt"
	"golang.org/x/net/context"
)

// Timezones map
type Timezones []map[string][]string

// GetTimezonesList gets the list of supported timezones
// http://docs.themoviedb.apiary.io/#reference/timezones/timezoneslist/get
func (tmdb *TMDb) GetTimezonesList(ctx context.Context) (*Timezones, error) {
	var timezoneList Timezones
	uri := fmt.Sprintf("%s/timezones/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(ctx, uri, &timezoneList)
	return result.(*Timezones), err
}
