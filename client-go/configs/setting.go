package configs

import "github.com/gin-gonic/gin"

func Setting(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "no configs."})
}
