package axios

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Asphaltt/promise"
)

// Get wraps http.Get as promise
func Get(url string) *promise.Promise {
	return promise.Init(http.Get(url))
}

// Post wraps http.Post as promise
func Post(url string, contentType string, data []byte) *promise.Promise {
	return promise.Init(http.Post(url, contentType, bytes.NewReader(data)))
}

// PostJSON wraps http.Post with json as promise
func PostJSON(url string, v interface{}) *promise.Promise {
	data, err := json.Marshal(v)
	if err != nil {
		return promise.Init(nil, err)
	}
	return Post(url, "application/json", data)
}
