module github.com/2637309949/bulrush-sso-example

go 1.13

require (
	github.com/2637309949/bulrush v0.0.0-20190904151601-2d79a67959e6
	github.com/2637309949/bulrush-addition v0.0.0-20190831034018-427428781eb0
	github.com/2637309949/bulrush-utils v0.0.0-20190815130414-1d9237d59ead
	github.com/gin-gonic/gin v1.4.0
	github.com/kataras/go-events v0.0.2
	github.com/michelloworld/ez-gin-template v0.0.0-20171028145326-7d77b0197797
)

// ## just for dev
replace github.com/2637309949/bulrush => ../../../bulrush

replace github.com/2637309949/bulrush-openapi => ../../../bulrush-openapi

replace github.com/2637309949/bulrush-addition => ../../../bulrush-addition

replace github.com/2637309949/bulrush-limit => ../../../bulrush-limit

replace github.com/2637309949/bulrush-template => ../../../bulrush-template

replace github.com/2637309949/bulrush-mq => ../../../bulrush-mq

replace github.com/2637309949/bulrush-role => ../../../bulrush-role

replace github.com/2637309949/bulrush-captcha => ../../../bulrush-captcha

replace github.com/2637309949/bulrush-delivery => ../../../bulrush-delivery

replace github.com/2637309949/bulrush-upload => ../../../bulrush-upload

replace github.com/2637309949/bulrush-logger => ../../../bulrush-logger

replace github.com/2637309949/bulrush-identify => ../../../bulrush-identify

replace github.com/2637309949/bulrush-proxy => ../../../bulrush-proxy

replace github.com/2637309949/bulrush-utils => ../../../bulrush-utils

replace github.com/2637309949/bulrush-sso => ../../bulrush-sso

// ## end