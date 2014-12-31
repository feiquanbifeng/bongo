// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import ()

var routes = bongo.NewRouter()

func init() {
    mc := &index.MainController{}
    routes.Get("/home", mc.FindById)
}
