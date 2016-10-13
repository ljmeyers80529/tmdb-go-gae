package tmdbgae

import (
	"fmt"
	"golang.org/x/net/context"
)

// Network struct
type Network struct {
	ID   int
	Name string
}

// GetNetworkInfo gets the basic information about a TV network
// http://docs.themoviedb.apiary.io/#reference/networks/networkid/get
func (tmdb *TMDb) GetNetworkInfo(ctx context.Context, id int) (*Network, error) {
	var networkInfo Network
	uri := fmt.Sprintf("%s/network/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(ctx, uri, &networkInfo)
	return result.(*Network), err
}
