package service

import (
	"context"
	"fmt"
	"snake-and-ladder/model/vo"
)

func RandomDice(ctx context.Context, board, player string) (int, error) {
	fmt.Println(board, player)
	return vo.NewDice(), nil
}
