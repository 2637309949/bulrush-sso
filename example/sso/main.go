// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"

	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
	eztemplate "github.com/michelloworld/ez-gin-template"
)

func main() {
	app := bulrush.Default()
	app.Config("cfg.yaml")
	app.Use(func(httpProxy *gin.Engine, router *gin.RouterGroup) {
		render := eztemplate.New()
		render.TemplatesDir = "../../template/"
		render.Layout = "layouts"
		render.Ext = ".html"
		render.Debug = false
		render.TemplateFuncMap = template.FuncMap{}
		httpProxy.HTMLRender = render.Init()
	})
	app.Use(func(event events.EventEmmiter) {
		event.On(bulrush.EventsRunning, func(message ...interface{}) {
			Logger.Info("EventsRunning %v", message)
		})
	})
	app.Run()
}
