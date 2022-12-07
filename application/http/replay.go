package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
	"net/http"
	"snake-and-ladder/service"
	"snake-and-ladder/utils"
)

func Replay(c *gin.Context) {
	board := c.MustGet("board").(string)
	player := c.MustGet("player").(string)
	ctx := c.Request.Context()
	replay, err := service.Replay(ctx, board, player)
	if err != nil {
		log.Errorf("Replay Gen: %v", err)
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.BaseRsp{
		Code: 0,
		Msg:  "Replay Gen Success",
		Body: replay,
	})
}
