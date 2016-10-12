package tmdbgae

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "strconv"
	"sync"
	"time"
)

const baseURL string = "https://api.themoviedb.org/3"

var (
	requestMutex   = &sync.Mutex{}
	rateLimitReset time.Time
)

// TMDb container struct for global properties
type TMDb struct {
	apiKey string
}

type apiStatus struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_message"`
}

// Init setup the apiKey
func Init(apiKey string) *TMDb {
	return &TMDb{apiKey: apiKey}
}
