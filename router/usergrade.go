package router

import (
	"github.com/gin-gonic/gin"

	"growth/api"
)

func InitUserGradeRouter(base *gin.RouterGroup) {
	base.GET("/ListGrades", api.GetListGrades)
}
