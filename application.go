// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sso

import (
	"github.com/gin-gonic/gin"
)

type (
	// Server defined Server side
	Server struct {
	}
	// Client defined Server side
	Client struct {
	}
)

// Plugin defined bulrush plugin
func (s *Server) Plugin(router *gin.RouterGroup) {

}

// Plugin defined bulrush plugin
func (c *Client) Plugin(router *gin.RouterGroup) {

}
