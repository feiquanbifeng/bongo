// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
    "net/http"
)

// Request serve
type Request struct {
    *http.Request
}
