package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
	"net/http"
	"snake-and-ladder/service"
	"snake-and-ladder/utils"
)

func RandomBoard(c *gin.Context) {
	ctx := c.Request.Context()
	board, err := service.RandomBoard(ctx)
	if err != nil {
		log.Errorf("Random Board: %v", err)
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, utils.BaseRsp{
			Code: 0,
			Msg:  "Random Board Success",
			Body: board,
		})
	}

}
