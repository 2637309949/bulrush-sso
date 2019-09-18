// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"html/template"
	"path"

	"github.com/thoas/go-funk"

	"github.com/2637309949/bulrush"
	delivery "github.com/2637309949/bulrush-delivery"
	identify "github.com/2637309949/bulrush-identify"
	sso "github.com/2637309949/bulrush-sso"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var users = []struct {
	UserName string
	Password string
	Email    string
}{
	struct {
		UserName string
		Password string
		Email    string
	}{
		UserName: "root",
		Password: "111111",
	},
}

// SSO server plugins
var SSO = (&sso.Server{}).Init(func(s *sso.Server) {
	// User Register
	s.Register = func(c *gin.Context, info struct {
		UserName string
		Password string
		Email    string
	}) error {
		one := funk.Find(users, func(u struct {
			UserName string
			Password string
			Email    string
		}) bool {
			return u.UserName == info.UserName && u.Password == info.Password
		})
		if one != nil {
			return errors.New("user has existed")
		}
		users = append(users, info)
		return nil
	}
})

// Delivery Upload, Logger, Captcha Plugin init
var Delivery = delivery.New().Init(func(d *delivery.Delivery) {
	d.URLPrefix = "/assets"
	d.Path = path.Join("../../public/", "")
})

// Identify plugin init
var Identify = identify.
	New().
	AddOptions(identify.FakeURLsOption([]string{`^/api/ignore$`, `^/api/gorm/mock`})).
	AddOptions(identify.FakeTokensOption([]string{})).
	AddOptions(identify.ModelOption(&identify.RedisModel{
		Redis: Redis,
	})).
	Init(func(iden *identify.Identify) {
		iden.AddOptions(
			identify.AuthOption(func(ctx *gin.Context) (interface{}, error) {
				var err error
				form := struct {
					UserName string `form:"username" json:"username" xml:"username" binding:"required"`
					Password string `form:"password" json:"password" xml:"password" binding:"required"`
				}{}
				if err = ctx.ShouldBind(&form); err != nil {
					return nil, err
				}
				one := funk.Find(users, func(u struct {
					UserName string
					Password string
					Email    string
				}) bool {
					return u.UserName == form.UserName && u.Password == form.Password
				})
				if one != nil {
					return map[string]string{
						"username": form.UserName,
					}, nil
				}
				return errors.New("User does not exist or password does not match"), nil
			}),
		)
	})

func addPlugin(app bulrush.Bulrush) bulrush.Bulrush {
	app.Use(Delivery)
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
	app.Use(Identify)
	app.Use(SSO)
	return app
}
