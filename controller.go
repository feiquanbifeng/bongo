// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "encoding/json"
    "html/template"
    "net/http"

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

func (c *Controller) ServeJson() {
    c.Ctx.Response.Out.Header("Content-Type", "application/json;charset=UTF-8")
    var (
        content []byte
        err     error
        data    = c.Data["json"]
    )
    if hasIndent {
        content, err = json.MarshalIndent(data, "", "  ")
    } else {
        content, err = json.Marshal(data)
    }
    if err != nil {
        http.Error(c.Ctx.Response.Out, err.Error(), http.StatusInternalServerError)
        return err
    }
    /*if coding {
        content = []byte(stringsToJson(string(content)))
    }*/
    c.Ctx.Request.Body(content)
    return nil
}

func (c *Controller) ServeXML() {

}

func (c *Controller) Redirect() {

}

func (c *Controller) Init(ctx *context.Context) error {
    c.Ctx = ctx
    return nil
}
