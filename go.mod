module github.com/2637309949/bulrush-sso

go 1.13

require (
	github.com/2637309949/bulrush-addition v0.0.0-20190923101740-7c6aa383b07f // indirect
	github.com/2637309949/bulrush-identify v0.0.0-20190911151611-398033f71556
	github.com/2637309949/bulrush-utils v0.0.0-20190831033838-023613c5526f
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.5+incompatible // indirect
	golang.org/x/sys v0.0.0-20190922100055-0a153f010e69 // indirect
)

// ## just for dev
replace github.com/2637309949/bulrush => ../bulrush

replace github.com/2637309949/bulrush-openapi => ../bulrush-openapi

replace github.com/2637309949/bulrush-addition => ../bulrush-addition

replace github.com/2637309949/bulrush-limit => ../bulrush-limit

replace github.com/2637309949/bulrush-template => ../bulrush-template

replace github.com/2637309949/bulrush-mq => ../bulrush-mq

replace github.com/2637309949/bulrush-role => ../bulrush-role

replace github.com/2637309949/bulrush-captcha => ../bulrush-captcha

replace github.com/2637309949/bulrush-delivery => ../bulrush-delivery

replace github.com/2637309949/bulrush-upload => ../bulrush-upload

replace github.com/2637309949/bulrush-logger => ../bulrush-logger

replace github.com/2637309949/bulrush-identify => ../bulrush-identify

replace github.com/2637309949/bulrush-proxy => ../bulrush-proxy

replace github.com/2637309949/bulrush-utils => ../bulrush-utils

// ## end
