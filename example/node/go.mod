module github.com/2637309949/bulrush-sso-example

go 1.13

require (
	github.com/2637309949/bulrush v0.0.0-20190904151601-2d79a67959e6
	github.com/2637309949/bulrush-addition v0.0.0-20190831034018-427428781eb0
	github.com/2637309949/bulrush-delivery v0.0.0-20190805055946-c208fdca9d47
	github.com/2637309949/bulrush-identify v0.0.0-20190831033932-28c237c53adf
	github.com/2637309949/bulrush-sso v0.0.0-20190911051105-087e987772de
	github.com/2637309949/bulrush-template v0.0.0-20190831034000-f9696cde578d
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.2+incompatible
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

replace github.com/2637309949/bulrush-sso => ../../../bulrush-sso

// ## end
