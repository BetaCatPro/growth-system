package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"growth/global"
	"growth/pb"
)

func GetListGrades(ctx *gin.Context) {
	out, err := global.ClientGrade.ListGrades(ctx, &pb.ListGradesRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    2,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, out)
}
