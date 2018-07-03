# ES6 Promsie in Go
Write functions like ES6 Promise.

`promise.go` is simple, which is less then 100 lines. Please read it.

## Example
API likes axios.js to do http request.

`ajax/ajax.go`
```
package ajax

import (
	"net/http"

	"github.com/Asphaltt/promise"
)

// Get wraps http.Get as promise
func Get(url string) *promise.Promise {
	return promise.Init(http.Get(url))
}
```

`examples/axios/main.go`
```
package main

import (
	"fmt"
	"net/http"

	"github.com/Asphaltt/promise/axios"
)

func main() {
	axios.Get("https://www.bing.com").Then(func(v interface{}) (interface{}, error) {
		resp := v.(*http.Response)
		fmt.Println(resp.Status)
		return nil, nil
	}).Catch(func(e error) {
		fmt.Println(e)
	})
}
```

## Another way
Another way to do axios with promise.

```
func main() {
    promise.Init(http.Get("https://www.bing.com")).Register(func(v interface{}) (interface{}, error) {
        resp := v.(*http.Response)
        fmt.Println(resp.Status)
        return nil, nil
    }).Catcher(func(e error) {
        fmt.Println(e)
    }).Done()
}
```