package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
	"net/http"
	"snake-and-ladder/service"
	"snake-and-ladder/utils"
)

func RandomDice(c *gin.Context) {
	board := c.MustGet("board").(string)
	player := c.MustGet("player").(string)
	ctx := c.Request.Context()
	dice, err := service.RandomDice(ctx, board, player)
	if err != nil {
		log.Errorf("Random Dice: %v", err)
		c.JSON(http.StatusBadRequest, utils.BaseRsp{
			Code: 1,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, utils.BaseRsp{
			Code: 0,
			Msg:  "Random Dice Success",
			Body: dice,
		})
	}

}
