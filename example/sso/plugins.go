// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"path"

	"github.com/2637309949/bulrush"
	delivery "github.com/2637309949/bulrush-delivery"
	sso "github.com/2637309949/bulrush-sso"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

// SSO server plugins
var SSO = &sso.Server{}

// Delivery Upload, Logger, Captcha Plugin init
var Delivery = delivery.New().Init(func(d *delivery.Delivery) {
	d.URLPrefix = "/assets"
	d.Path = path.Join("../../public/", "")
})

func addPlugin(app bulrush.Bulrush) bulrush.Bulrush {
	app.Use(func(httpProxy *gin.Engine, router *gin.RouterGroup) {
		httpProxy.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
			Root:         "../../template",
			Extension:    ".html",
			Master:       "layout",
			Partials:     []string{},
			Funcs:        template.FuncMap{},
			DisableCache: true,
		})
	})
	app.Use(SSO)
	app.Use(Delivery)
	return app
}
