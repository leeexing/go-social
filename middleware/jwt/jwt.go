package jwt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leeexing/go-social/pkg/e"
	"github.com/leeexing/go-social/pkg/util"
)

// JWT 校验Token中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			code = e.ERROR_AUTH_HEADERS
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			code = e.ERROR_AUTH_FORMAT
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}

		token := parts[1]
		// 简单的权限字段验证
		// claims, err := util.ParseToken(token)
		// 验证非本服务发票的token
		claims, err := util.ParseNtToken(token)
		fmt.Printf("claims: %#v \n", claims)
		if err != nil {
			fmt.Printf("Err: %#v, type: %T \n", err, err)
			log.Println(err)
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			data = fmt.Sprintf("%s", err)
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		// 拿到 claims 可以进一步判断里面的 Audience 字段是否符合配置要求
		// code

		// 这里可以直接将token里面有用的信息放到请求的上下文
		c.Set("username", claims.NtName)
		c.Set("useid", claims.NtUID)

		c.Next()
	}
}
