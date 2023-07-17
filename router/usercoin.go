package router

import (
	"github.com/gin-gonic/gin"

	"growth/api"
)

func InitUserCoinRouter(base *gin.RouterGroup) {
	base.GET("/ListTasks", api.GetListTasks)
	base.POST("/UserCoinChange", api.PostUserCoinChange)
}
