// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package index

import (
    "io"
    "net/http"

    "bongo"
)

type MainController struct {
    bongo.Controller
}

func (m *MainController) FindById(id int) {
    m.Data["json"] = `{"Name": "Hello"}`
    m.ServeJSON()
}
