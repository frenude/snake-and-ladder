package service

import (
	"context"
	"snake-and-ladder/dao/redis"
	"snake-and-ladder/model/vo"
)

func RandomDice(ctx context.Context, board, player string) (int, error) {

	randomInt := vo.NewDice()
	err := redis.SetDice(ctx, board, player, randomInt, 0, 0, 0, false)
	if err != nil {
		return 0, err
	}
	return randomInt, nil
}
