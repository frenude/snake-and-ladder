package service

import (
	"context"
	"encoding/json"
	"snake-and-ladder/dao/db"
	"snake-and-ladder/dao/redis"
	"snake-and-ladder/model/po"
	"snake-and-ladder/model/vo"
	"strconv"
)

func RandomBoard(ctx context.Context, n int) (*vo.Board, error) {
	snake, ladder := vo.RandomSnakeAndLadder()
	marshal_snake, err := json.Marshal(snake)
	marshal_ladder, err := json.Marshal(ladder)
	if err != nil {
		return nil, err
	}
	b := po.Board{
		Snake:  string(marshal_snake),
		Ladder: string(marshal_ladder),
	}
	err = db.CreateBoard(ctx, &b)
	if err != nil {
		return nil, err
	}
	board, err := vo.NewBoard(n, strconv.Itoa(int(b.Id)), snake, ladder)
	if err != nil {
		return nil, err
	}
	err = redis.SetBoard(ctx, board)
	if err != nil {
		return nil, err
	}
	return board, nil
}
