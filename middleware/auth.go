package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_template/log"
	"web_template/utils"
)

// AuthMiddleware 从body或url query中获取token，验证token是否有效
func AuthMiddleware(ctx *gin.Context) {
	var token string
	reqAuth := ctx.Request.Header.Get("Authorization")

	if reqAuth == "" {
		// 从query中获取token
		token = ctx.Query("token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "no Authorization field",
			})
			ctx.Abort()
			return
		}
	} else {
		// 从body中获取token
		s := bytes.Split([]byte(reqAuth), []byte(" "))
		if len(s) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "no Authorization field",
			})
			ctx.Abort()
			return
		}
		token = string(s[1])
	}

	ok, err := utils.ValidToken(token)
	if err != nil {
		log.GetLogger().Debug("token error: ", err)
		ctx.Abort()
		return
	}
	if !ok {
		log.GetLogger().Debug("token invalid: ", token)
		ctx.Abort()
		return
	}

	ctx.Next()
}
