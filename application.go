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
		Register func(*gin.Context, struct {
			UserName string
			Password string
			Email    string
		}) error
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
func (s *Server) Plugin(router *gin.RouterGroup, iden *identify.Identify) {
	// skip identify check
	*iden.FakeURLs = append(*iden.FakeURLs, utils.Some(s.Route["home"], "/").(string))
	*iden.FakeURLs = append(*iden.FakeURLs, utils.Some(s.Route["login"], "/login").(string))
	*iden.FakeURLs = append(*iden.FakeURLs, utils.Some(s.Route["logout"], "/logout").(string))
	*iden.FakeURLs = append(*iden.FakeURLs, utils.Some(s.Route["register"], "/register").(string))

	router.GET(utils.Some(s.Route["home"], "/").(string), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", "")
	})

	// handle for login
	router.GET(utils.Some(s.Route["login"], "/login").(string), func(c *gin.Context) {
		query := c.Request.URL.RawQuery
		_csrf, _ := c.Get("csrf")
		token := iden.GetToken(c)
		if token != nil {
			valid := iden.VerifyToken(token.AccessToken)
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
		var message string
		ret, err := iden.Auth(c)
		_csrf, _ := c.Get("csrf")
		query := c.Request.URL.RawQuery
		if err == nil && ret != nil {
			var info *identify.Token
			info, err = iden.ObtainToken(ret)
			if err == nil && info != nil {
				c.SetCookie("accessToken", info.AccessToken, 60*60*int(iden.ExpiresIn), "/", "", false, true)
				redirect := utils.Some(c.Query("redirect_uri"), c.Query("redirect")).(string)
				if redirect != "" {
					if strings.Contains(redirect, "?") {
						c.Redirect(http.StatusMovedPermanently, redirect+"&accessToken="+info.AccessToken)
					} else {
						c.Redirect(http.StatusMovedPermanently, redirect+"?accessToken="+info.AccessToken)
					}
					return
				}
				c.Redirect(http.StatusMovedPermanently, s.Route["home"])
				return
			}
		}
		if err != nil && message == "" {
			message = err.Error()
		} else {
			message = "User does not exist or password does not match"
		}
		c.HTML(http.StatusOK, "pages/login", map[string]interface{}{"_csrf": _csrf, "action": s.Route["login"] + "?" + query, "message": message})
	})

	// logou route
	router.GET(utils.Some(s.Route["logout"], "/logout").(string), func(c *gin.Context) {
		token := iden.GetToken(c)
		iden.Model.Revoke(token)
		c.SetCookie("accessToken", "", 0, "/", "", false, true)
		redirect := utils.Some(c.Query("redirect_uri"), c.Query("redirect")).(string)
		redirect = utils.Some(redirect, s.Route["home"]).(string)
		c.Redirect(http.StatusMovedPermanently, redirect)
	})

	// register route
	router.POST(utils.Some(s.Route["register"], "/register").(string), func(c *gin.Context) {
		var err error
		var message string
		_csrf, _ := c.Get("csrf")
		query := c.Request.URL.RawQuery
		form := struct {
			UserName string `form:"username" json:"username" xml:"username"`
			Password string `form:"password" json:"password" xml:"password"`
			Email    string `form:"email" json:"email" xml:"email"`
		}{}
		if err = c.ShouldBind(&form); err == nil {
			err = s.Register(c, struct {
				UserName string
				Password string
				Email    string
			}{
				UserName: form.UserName,
				Password: form.Password,
				Email:    form.Email,
			})
			if err == nil {
				c.HTML(http.StatusOK, "pages/login", map[string]interface{}{"_csrf": _csrf, "action": s.Route["login"] + "?" + query})
			}
		} else {
			if err != nil && message == "" {
				message = err.Error()
			} else {
				message = "Incorrect username or password"
			}
			c.HTML(http.StatusOK, "pages/login", map[string]interface{}{"_csrf": _csrf, "action": s.Route["login"] + "?" + query, "message": message})
		}
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
