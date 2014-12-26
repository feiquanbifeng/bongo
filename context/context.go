// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import ()

// Http Request context struct
// Input and Output provider flexible api
type Context struct {
    *Request
    *Response
}
