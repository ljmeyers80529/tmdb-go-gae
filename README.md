# tmdb_go_gae

Based on code by ryanbradynd05, modified to be used with the google app engine.
[GoDoc](https://godoc.org/github.com/ljmeyers80529/tmdb-go-gae)
=================================

A Go Wrapper for the API of [The Movie DB](http://www.themoviedb.org/). Complete [documentation](https://godoc.org/github.com/ljmeyers80529/tmdb-go-gae)

An **api_key** is needed to use the API. Register for one at [themoviedb.org](https://www.themoviedb.org/documentation/api).

Note: This product uses the TMDb API but is not endorsed or certified by TMDb.

<img src="https://assets.tmdb.org/images/logos/var_7_0_tmdb-logo-2_Antitled.svg" alt="The Movie DB" width="200" height="200" />

## How to install

```shell
go get github.com/ryanbradynd05/go-tmdb
```

## How to use

Import the libraries

```go
import "github.com/ljmeyers80529/tmdb-go-gae"
import "google.golang.org/appengine"
```


Initialize the library using your api_key
```go
func Init(apiKey string) *TMDb {
  return &TMDb{apiKey: apiKey}
}
```

Use the api methods as you want, for example:

```go
ctx := appengine.NewContext(req)
fightClubInfo, err := TMDb.GetMovieInfo(ctx, 550, nil)
```

To use optional parameters, pass in a map[string]string of options and values:

```go
var options = make(map[string]string)
options["language"] = "es"
spanishFightClub, err := TMDb.GetMovieInfo(ctx, 550, options)
```

All functions return Go structs. To return JSON, use the ToJSON function:

```go
fightClubJson, err := TMDb.ToJSON(fightClubInfo)
```
