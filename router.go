// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "log"
    "net/http"
    "path"

    "bongo/context"

    "github.com/gorilla/mux"
)

var RR *Routes

func init() {
    RR = &Routes{
        Controller: make(map[string]ControllerInterface),
        R:          mux.NewRouter(),
    }
}

type HttpApiFunc func(args ...interface{}) error

type Router interface {
    Get(string, HttpApiFunc)
    Post(string, HttpApiFunc)
}

type Routes struct {
    path       string
    Controller map[string]ControllerInterface
    R          *mux.Router
}

func newRouter(r string, fn func(...interface{}) error) *Routes {
    RR.path = r
    fn()
    return RR
}

// Wrap http handler
func makeHttpHandler(path string, localMethod string, localRoute string, handlerFunc HttpApiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // log the request
        context := &context.Context{
            Response: &context.Response{Out: w},
            Request:  &context.Request{r},
        }
        c := RR.Controller[path]
        c.Init(context)
        r.ParseForm()
        if err := handlerFunc(r.Form); err != nil {
            log.Fatalf("Handler for %s %s returned error: %s", localMethod, localRoute, err)
        }
    }
}

// Create route func for private
func (r *Routes) createRoute(m string, route string, h HttpApiFunc) {
    f := makeHttpHandler(r.path, m, route, h)
    r.R.Path(path.Clean(r.path + route)).Methods(m).HandlerFunc(f)
}

// Get request
func (r *Routes) Get(route string, h HttpApiFunc) {
    r.createRoute("GET", route, h)
}

// Post request
func (r *Routes) Post(route string, h HttpApiFunc) {
    r.createRoute("POST", route, h)
}

func (r *Routes) RegisterController(c ControllerInterface) error {
    RR.Controller[r.path] = c
    return nil
}
