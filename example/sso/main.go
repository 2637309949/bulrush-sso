// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/2637309949/bulrush"
	"github.com/kataras/go-events"
)

func main() {
	app := bulrush.Default()
	app.Config("cfg.yaml")
	app = addPlugin(app)
	app.Use(func(event events.EventEmmiter) {
		event.On(bulrush.EventsRunning, func(message ...interface{}) {
			Logger.Info("EventsRunning %v", message)
		})
	})
	app.Run()
}
