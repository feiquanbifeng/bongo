// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"bongo/router"
	"mvc/controllers/index"
)

func Init() {
	var r = router.Routes
	r.Get("/home", index.index)
	r.Post("/abc", index.hello)
}
