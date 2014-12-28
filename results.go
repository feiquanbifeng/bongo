// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "bongo/context"
)

type Result interface {
    Apply(req *context.Request, resp *context.Response)
}
