// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "net/http"
    "path"
    "reflect"
    "strconv"
    "time"

    "github.com/codegangsta/inject"
    "github.com/gorilla/mux"
)

// Define type interface accept different route
type Handler interface{}

// Validate handler if it's not a func type
func validateHandler(handler Handler) {
    if reflect.TypeOf(handler).Kind() != reflect.Func {
        panic("bongo handler must be a callable func")
    }
}

// Define Bongo struct
type Bongo struct {
    inject.Injector
    *Logger
    handlers []Handler
    config   map[string]interface{}
}

// New a Bongo application
func NewBongo() *Bongo {
    b := &Bongo{
        Injector: inject.New(),
        Logger:   NewLogger(),
    }
    return b
}

//  Application include Bongo instance and router information
//  package main
//
//  import "github.com/feiquanbifeng/bongo"
//
//  func main() {
//    b := bongo.App()
//
//    b.Get("/", func() string {
//      return "Hello world!"
//    })
//
//    b.Run(3000)
//  }
type Application struct {
    *Bongo
    Router
}

func App() *Application {
    r := NewRouter()
    b := NewBongo()
    b.MapTo(r, (*Routes)(nil))
    b.Action(r.Handle)
    return &Application{b, r}
}

// Run the server
// If port is provider use it
func (b *Bongo) Run(port int) error {
    var addr string
    if port == 0 {
        addr = ":3000"
    } else {
        addr = ":" + strconv.Itoa(port)
    }
    chErrors := make(chan error, 1)
    go func() {
        chErrors <- b.listen(addr)
    }()
    return <-chErrors
}

// ServeHTTP is the HTTP Entry point for a Bongo instance.
// Useful if you want to control your own HTTP server.
func (b *Bongo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // Clean path to canonical form and redirect.
    if p := cleanPath(req.URL.Path); p != req.URL.Path {

        // Added 3 lines (Philip Schlump) - It was droping the query string and #whatever from query.
        // This matches with fix in go 1.2 r.c. 4 for same problem.  Go Issue:
        // http://code.google.com/p/go/issues/detail?id=5252
        url := *req.URL
        url.Path = p
        p = url.String()

        w.Header().Set("Location", p)
        w.WriteHeader(http.StatusMovedPermanently)
        return
    }
}

// Start the app
func (b *Bongo) listen(port string) error {
    server := http.Server{
        Addr:        port,
        Handler:     b,
        ReadTimeout: 5 * time.Second,
    }
    b.Info("listen on %s (%s)\n", port, "dev")
    return server.ListenAndServe()
}

// Get the value from the config variable
func (b *Bongo) Get(key string) interface{} {
    return b.config[key]
}

// Set the value for config variable or with pair
// If have the same value the before value will be override
func (b *Bongo) Set(key string, value interface{}) {
    if b.config == nil {
        b.config = make(map[string]interface{})
    }
    b.config[key] = value
}

// Use adds a middleware Handler to the stack.
// Will panic if the handler is not a callable func.
// Middleware Handlers are invoked in the order that they are added.
func (b *Bongo) Use(handler Handler) {
    validateHandler(handler)
    b.handlers = append(b.handlers, handler)
}

// cleanPath returns the canonical path for p, eliminating . and .. elements.
// Borrowed from the net/http package.
func cleanPath(p string) string {
    if p == "" {
        return "/"
    }
    if p[0] != '/' {
        p = "/" + p
    }
    np := path.Clean(p)
    // path.Clean removes trailing slash except for root;
    // put the trailing slash back if necessary.
    if p[len(p)-1] == '/' && np != "/" {
        np += "/"
    }
    return np
}
