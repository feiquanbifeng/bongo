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

// Router registers Routes to be matched and dispatches a handler.
// This will send all incoming requests to the router.
type Router interface {
    Get(string, Handler)
    Post(string, Handler)
}

// NewRouter returns a new router instance.
func NewRouter() Router {

}

type Route struct {
    path       string
    Controller map[string]ControllerInterface
    R          *mux.Router
}

// Wrap http handler
func makeHttpHandler(path string, localMethod string, localRoute string, handlerFunc Handler) http.HandlerFunc {
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
func (r *Route) createRoute(m string, route string, h Handler) {
    f := makeHttpHandler(r.path, m, route, h)
    r.R.Path(path.Clean(r.path + route)).Methods(m).HandlerFunc(f)
}

// Get request
func (r *Route) Get(route string, h Handler) {
    r.createRoute("GET", route, h)
}

// Post request
func (r *Route) Post(route string, h Handler) {
    r.createRoute("POST", route, h)
}

func (r *Route) RegisterController(c ControllerInterface) error {
    RR.Controller[r.path] = c
    return nil
}
