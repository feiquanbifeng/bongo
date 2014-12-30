// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "net/http"
    "reflect"
    "strconv"
    "time"

    "github.com/codegangsta/inject"
    "github.com/gorilla/mux"
)

// Define type interface accept different route
type Handler interface{}

func validateHandler(handler Handler) {
    if reflect.TypeOf(handler).Kind() != reflect.Func {
        panic("bongo handler must be a callable func")
    }
}

// Define Bongo struct
type Bongo struct {
    inject.Injector
    *Logger
    handler *mux.Router
    config  map[string]interface{}
}

// New a bongo application
func NewBongo() *Bongo {
    b := &Bongo{
        Injector: inject.New(),
        Logger:   NewLogger(),
    }
    return b
}

// Application include Bongo instance and router information
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
        addr = ":8080"
    } else {
        addr = ":" + strconv.Itoa(port)
    }
    chErrors := make(chan error, 1)
    go func() {
        chErrors <- b.listen(addr)
    }()

    err := <-chErrors
    if err != nil {
        return err
    }
    return nil
}

func Route(route string, fn func()) error {
    h := newRouter(route, fn)
    bon.handler = h.R
    return nil
}

// Start the app
func (b *Bongo) listen(port string) error {
    server := http.Server{
        Addr:        port,
        Handler:     b.handler,
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
// If have the same value the before value will be over
func (b *Bongo) Set(key string, value interface{}) {
    if b.config == nil {
        b.config = make(map[string]interface{})
    }
    b.config[key] = value
}
