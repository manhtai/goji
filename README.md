# About Goji

[![Test](https://github.com/manhtai/goji/actions/workflows/test.yml/badge.svg)](https://github.com/manhtai/goji/actions/workflows/test.yml)

Goji is a HTTP request multiplexer, similar to [`net/http.ServeMux`][servemux].
It compares incoming requests to a list of registered [Patterns][pattern], and
dispatches to the [http.Handler][handler] that corresponds to the first
matching Pattern. Goji also supports [Middleware][middleware] (composable
shared functionality applied to every request) and uses the standard
[`context`][context] package to store request-scoped values.

[servemux]: https://golang.org/pkg/net/http/#ServeMux
[pattern]: https://godoc.org/goji.io#Pattern
[handler]: https://golang.org/pkg/net/http/#Handler
[middleware]: https://godoc.org/goji.io#Mux.Use
[context]: https://golang.org/pkg/context

## Quick Start

```go
package main

import (
        "fmt"
        "net/http"

        "github.com/kenshaw/goji"
)

func hello(w http.ResponseWriter, r *http.Request) {
        name := pat.Param(r, "name")
        fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
        mux := goji.NewMux()
        mux.HandleFunc(pat.Get("/hello/:name"), hello)

        http.ListenAndServe("localhost:8000", mux)
}
```

Please refer to [Goji's GoDoc Documentation][godoc] for a full API reference.

[godoc]: https://godoc.org/github.com/kenshaw/goji
