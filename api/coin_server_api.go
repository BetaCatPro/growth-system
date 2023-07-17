package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"growth/global"
	"growth/pb"
)

func GetListTasks(ctx *gin.Context) {
	out, err := global.ClientCoin.ListTasks(ctx, &pb.ListTasksRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    2,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, out)
}

func PostUserCoinChange(ctx *gin.Context) {
	body := &pb.UserCoinChangeRequest{}
	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    2,
			"message": err.Error(),
		})
	} else if out, err := global.ClientCoin.UserCoinChange(ctx, body); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    2,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, out)
	}
	ctx.JSON(http.StatusOK, nil)
}
