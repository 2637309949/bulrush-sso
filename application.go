// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sso

import (
	"net/http"
	"strings"

	identify "github.com/2637309949/bulrush-identify"
	utils "github.com/2637309949/bulrush-utils"
	"github.com/gin-gonic/gin"
)

type (
	// Config for Client and Server
	Config struct {
		Route    map[string]string
		Pages    map[string]string
		Binding  func(*gin.Context, interface{})
		Register func(*gin.Context, interface{})
	}
	// Server defined Server side
	Server struct {
		Config
	}
	// Client defined Server side
	Client struct {
		Config
	}
)

// Init s
func (s *Server) Init(init func(*Server)) *Server {
	init(s)
	return s
}

// Plugin defined bulrush plugin
func (s *Server) Plugin(router *gin.RouterGroup, identify *identify.Identify) {
	// skip identify check
	*identify.FakeURLs = append(*identify.FakeURLs, utils.Some(s.Route["home"], "/").(string))
	*identify.FakeURLs = append(*identify.FakeURLs, utils.Some(s.Route["login"], "/login").(string))
	*identify.FakeURLs = append(*identify.FakeURLs, utils.Some(s.Route["logout"], "/logout").(string))
	*identify.FakeURLs = append(*identify.FakeURLs, utils.Some(s.Route["register"], "/register").(string))

	router.GET(utils.Some(s.Route["home"], "/").(string), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", "")
	})

	// handle for login
	router.GET(utils.Some(s.Route["login"], "/login").(string), func(c *gin.Context) {
		query := c.Request.URL.RawQuery
		_csrf, _ := c.Get("csrf")
		token := identify.GetToken(c)
		if token != nil {
			valid := identify.VerifyToken(token.AccessToken)
			if !valid {
				c.HTML(http.StatusOK, "pages/login", map[string]interface{}{"_csrf": _csrf, "action": s.Route["login"] + "?" + query})
			} else {
				redirect := utils.Some(c.Query("redirect_uri"), c.Query("redirect")).(string)
				if redirect != "" {
					if strings.Contains(redirect, "?") {
						c.Redirect(http.StatusMovedPermanently, redirect+"&accessToken="+token.AccessToken)
					} else {
						c.Redirect(http.StatusMovedPermanently, redirect+"?accessToken="+token.AccessToken)
					}
				} else {
					c.Redirect(http.StatusMovedPermanently, s.Route["home"])
				}
			}
		} else {
			c.HTML(http.StatusOK, "pages/login", map[string]interface{}{"_csrf": _csrf, "action": s.Route["login"] + "?" + query})
		}
	})

	// post login
	router.POST(utils.Some(s.Route["login"], "/login").(string), func(c *gin.Context) {
	})
	router.GET(utils.Some(s.Route["logout"], "/logout").(string), func(c *gin.Context) {
	})
	router.POST(utils.Some(s.Route["register"], "/register").(string), func(c *gin.Context) {
	})
}

// Init c
func (c *Client) Init(init func(*Client)) *Client {
	init(c)
	return c
}

// Plugin defined bulrush plugin
func (c *Client) Plugin(router *gin.RouterGroup) {
	router.GET(utils.Some(c.Route["login"], "/login").(string), func(c *gin.Context) {
	})
	router.GET(utils.Some(c.Route["logout"], "/logout").(string), func(c *gin.Context) {
	})
	router.GET(utils.Some(c.Route["redirect"], "/redirect").(string), func(c *gin.Context) {
	})
}
