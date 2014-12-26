// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bongo

import (
    "net/http"
    "strconv"
    "time"

    "github.com/codegangsta/inject"
)

// Define type interface accept different route
type Handler interface{}

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

type Bongo struct {
    inject.Injector
    port   string
    log    *Logger
    config map[string]interface{}
}

func NewBongo() *Bongo {
    b := &Bongo{
        Injector: inject.New(),
        log:      newLogger(),
    }
    return b
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
    b.port = addr
    chErrors := make(chan error, 1)
    go func() {
        chErrors <- b.listen()
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

// start the app
func (b *Bongo) listen() error {
    server := http.Server{
        Addr:        b.port,
        Handler:     b.handler,
        ReadTimeout: 5 * time.Second,
    }
    return server.ListenAndServe()
}

// Get the value from the config variable
func (b *Bongo) Get(key string) interface{} {
    return b.config[key]
}

// Set the value for config variable or with pair
func (b *Bongo) Set(key string, values ...interface{}) {
    if b.config == nil {
        b.config = make(map[string]interface{})
    }

    if len(values) == 0 {
        return b.settings[settings]
    } else {
        // get the first value
        for _, v := range args {
            b.settings[settings] = v
            break
        }
    }
}
