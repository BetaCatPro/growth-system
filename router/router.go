package router

import (
	"github.com/gin-gonic/gin"

	"growth/middleware"
)

func GetRouters() *gin.Engine {
	router := gin.New()

	// 用户积分服务的方法
	v1Group := router.Group("/v1", middleware.Cors())
	gUserCoin := v1Group.Group("/Growth.UserCoin")

	// 用户等级服务的方法
	gUserGrade := v1Group.Group("/Growth.UserGrade")

	InitUserCoinRouter(gUserCoin)
	InitUserGradeRouter(gUserGrade)

	return router
}
