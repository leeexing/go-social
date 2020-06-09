package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/leeexing/go-social/pkg/setting"
	"github.com/leeexing/go-social/routers/api"
	"github.com/leeexing/go-social/middleware/jwt"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Static("/static", "./static")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "good",
				"data": nil,
			})
		})
	}

	return r
}
