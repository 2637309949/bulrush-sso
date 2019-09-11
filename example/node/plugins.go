// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/2637309949/bulrush"
	identify "github.com/2637309949/bulrush-identify"
	sso "github.com/2637309949/bulrush-sso"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-gonic/gin"
)

// SSO server plugins
var SSO = &sso.Client{}

// Identify plugin init
var Identify = identify.
	New().
	AddOptions(identify.FakeURLsOption([]string{`^/api/ignore$`, `^/api/gorm/mock`})).
	AddOptions(identify.FakeTokensOption([]string{})).
	AddOptions(identify.ModelOption(&identify.RedisModel{
		Redis: addition.Redis,
	})).
	Init(func(iden *identify.Identify) {
		iden.AddOptions(
			identify.AuthOption(func(ctx *gin.Context) (interface{}, error) {
				var user interface{}
				form := struct {
					UserName string `form:"username" json:"username" xml:"username" binding:"required"`
					Password string `form:"password" json:"password" xml:"password" binding:"required"`
				}{}
				if err := ctx.ShouldBind(&form); err != nil {
					return nil, err
				}
				if form.UserName == "root" && form.Password == "111111" {
					return map[string]string{
						"username": form.UserName,
					}, nil
				}
				return user, nil
			}),
		)
	})

func addPlugin(app bulrush.Bulrush) bulrush.Bulrush {
	app.Use(SSO)
	app.Use(Identify)
	return app
}
