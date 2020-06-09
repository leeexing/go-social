module github.com/leeexing/go-social

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/jinzhu/gorm v1.9.12
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.57.0 // indirect
)

replace (
	github.com/leeexing/go-social/conf => E:/Go/src/nuctech.com/go-social/conf
	github.com/leeexing/go-social/middleware => E:/Go/src/nuctech.com/go-social/middleware
	github.com/leeexing/go-social/models => E:/Go/src/nuctech.com/go-social/models
	github.com/leeexing/go-social/pkg/e => E:/Go/src/nuctech.com/go-social/pkg/e
	github.com/leeexing/go-social/pkg/logging => E:/Go/src/nuctech.com/go-social/pkg/logging
	github.com/leeexing/go-social/pkg/setting => E:/Go/src/nuctech.com/go-social/pkg/setting
	github.com/leeexing/go-social/pkg/util => E:/Go/src/nuctech.com/go-social/pkg/util
	github.com/leeexing/go-social/routers => E:/Go/src/nuctech.com/go-social/routers
	github.com/leeexing/go-social/routers/api => E:/Go/src/nuctech.com/go-social/routers/api
)
