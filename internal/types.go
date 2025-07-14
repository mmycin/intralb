package internal

import "net/http"

type Request struct {
    W http.ResponseWriter
    R *http.Request
}
