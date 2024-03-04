package util

import "github.com/gin-gonic/gin"

func CORSHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method //获取请求包的http请求方法
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		ctx.Header("Access-Control-Max-Age", "172800")
		ctx.Header("Access-Control-Allow-Credentials", "false")
		// ctx.Set("content-type", "application/json") //设置返回格式是json，不能设置这个，否则bind绑定数据就为空了

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.JSON(200, "Options Request!")
		}

		ctx.Next()
	}
}
