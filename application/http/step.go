package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
	"net/http"
	"snake-and-ladder/model/vo"
	"snake-and-ladder/service"
	"snake-and-ladder/utils"
)

func Step(c *gin.Context) {
	board := c.MustGet("board").(string)
	player := c.MustGet("player").(string)
	ctx := c.Request.Context()
	chess := vo.Chess{}
	err := c.BindJSON(&chess)
	if chess.Move > 6 || chess.Move < 1 {
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  errors.New("move 参数输入范围错误 请输入投掷骰子点数").Error(),
		})
		return
	}
	if chess.Point > 100 || chess.Point < 0 {
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  errors.New("point 参数输入范围错误 请输入上一步位置或者0").Error(),
		})
		return
	}
	if err != nil {
		log.Errorf("Stem Json Parse Error: %v", err)
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	step, err := service.Step(ctx, board, player, &chess)
	if err != nil {
		log.Errorf("Step Gen Error: %v", err)
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.BaseRsp{
		Code: 0,
		Msg:  "Step  Gen Success",
		Body: step,
	})
}
