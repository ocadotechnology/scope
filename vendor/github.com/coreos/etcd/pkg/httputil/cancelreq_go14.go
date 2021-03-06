// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// borrowed from golang/net/context/ctxhttp/cancelreq_go14.go

// +build !go1.5

package httputil

import "net/http"

type requestCanceler interface {
	CancelRequest(req *http.Request)
}

func RequestCanceler(rt http.RoundTripper, req *http.Request) func() {
	c, ok := rt.(requestCanceler)
	if !ok {
		return func() {}
	}
	return func() {
		c.CancelRequest(req)
	}
}
