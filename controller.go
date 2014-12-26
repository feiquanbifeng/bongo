// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "html/template"
    // "net/http"

    "bongo/context"
)

type ControllerInterface interface {
    Render(tpl string, data interface{}, mime string) error
    Init(ctx *context.Context) error
}

type Controller struct {
    Ctx   *context.Context
    Route *Routes
}

func (c *Controller) Render(tpl string, data interface{}, mime string) error {
    ctx := c.Ctx.Response.Out
    ctx.Header().Set("Content-Type", "text/html; charset=utf-8")
    t, err := template.ParseFiles(tpl)
    if err != nil {
        return err
    }
    t.Execute(ctx, data)
    return nil
}

func (c *Controller) Redirect() {

}

func (c *Controller) Init(ctx *context.Context) error {
    c.Ctx = ctx
    return nil
}
