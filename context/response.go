// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
    "net/http"
)

type Status int

const (
    StatusOk       Status = 0
    StatusErr      Status = 1
    StatusNotFound Status = 404
)

// Response http serve
type Response struct {
    Status      int
    ContentType string

    Out http.ResponseWriter
}
