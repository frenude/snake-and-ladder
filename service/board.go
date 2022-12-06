package service

import (
	"context"
	"snake-and-ladder/dao/redis"
	"snake-and-ladder/model/vo"
)

func RandomBoard(ctx context.Context) (*vo.Board, error) {
	board, err := vo.NewBoard(4)
	if err != nil {
		return nil, err
	}
	err = redis.SetBoard(ctx, board)
	if err != nil {
		return nil, err
	}
	return board, nil
}
