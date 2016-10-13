package tmdbgae

import (
	"fmt"
	"golang.org/x/net/context"
)

// Company struct
type Company struct {
	Description   string
	Headquarters  string
	Homepage      string
	ID            int
	LogoPath      string `json:"logo_path"`
	Name          string
	Movies        *MoviePagedResults `json:",omitempty"`
	ParentCompany struct {
		Name     string
		ID       int
		LogoPath string `json:"logo_path"`
	} `json:"parent_company"`
}

// GetCompanyInfo gets all of the basic information about a company
// http://docs.themoviedb.apiary.io/#reference/companies/companyid/get
func (tmdb *TMDb) GetCompanyInfo(ctx context.Context, id int, options map[string]string) (*Company, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var companyInfo Company
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/company/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(ctx, uri, &companyInfo)
	return result.(*Company), err
}

// GetCompanyMovies gets the list of movies associated with a particular company
// http://docs.themoviedb.apiary.io/#reference/companies/companyidmovies/get
func (tmdb *TMDb) GetCompanyMovies(ctx context.Context, id int, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var movies MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/company/%v/movies?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(ctx, uri, &movies)
	return result.(*MoviePagedResults), err
}
