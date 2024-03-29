// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	addition "github.com/2637309949/bulrush-addition"
	"github.com/2637309949/bulrush-addition/logger"
	redisext "github.com/2637309949/bulrush-addition/redis"
	"github.com/go-redis/redis"
)

// Logger defined bulrush or system log global proxy
// Two Transport has been added
// 1: RotateFile
// 2: Console
var Logger = addition.RushLogger.
	AddTransports(
		&logger.Transport{
			Dirname: "error",
			Level:   logger.ERROR,
			Maxsize: logger.Maxsize,
		},
		&logger.Transport{
			Dirname: "combined",
			Level:   logger.SILLY,
			Maxsize: logger.Maxsize,
		},
	).
	Init(func(j *logger.Journal) {
		j.SetFlags((logger.LstdFlags | logger.Llongfile))
	})

// Redis application redis store
var Redis = redisext.New().
	Init(func(r *redisext.Redis) {
		r.Conf(&redis.Options{
			Addr: "127.0.0.1:6379",
		})
	})
