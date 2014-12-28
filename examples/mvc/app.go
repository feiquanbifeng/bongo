// Copyright 2014 The Bongo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
    "bongo"
    "log"
    "os"

    "mvc/controllers/index"
)

func main() {
    app := bongo.NewBongo()
    dirname, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }
    app.Set("views", dirname+"\\views")

    // serve static files
    // app.Use(bongo.Static(dirname + "/public"));

    // Register routes
    app.Get("/", index.FindById)
    // Run the app
    app.Run(":3000")
}
