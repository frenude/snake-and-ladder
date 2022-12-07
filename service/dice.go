package service

import (
	"context"
	"encoding/json"
	goredis "github.com/go-redis/redis/v8"
	"snake-and-ladder/dao/db"
	"snake-and-ladder/dao/redis"
	"snake-and-ladder/model/vo"
	"strconv"
)

func RandomDice(ctx context.Context, board, player string) (int, error) {
	_, err := redis.GetDice(ctx, board, player)
	if err == goredis.Nil {
		i, err := strconv.Atoi(board)
		if err != nil {
			return 0, err
		}
		begin, throw, err := db.SelectStep(ctx, uint(i), player)
		if begin != nil && throw != nil {
			getBoard, err := redis.GetBoard(ctx, board)
			var snake [][]int
			var ladder [][]int
			if err == goredis.Nil {
				selectBoard, err := db.SelectBoard(ctx, uint(i))
				if err != nil {
					return 0, err
				}
				err = json.Unmarshal([]byte(selectBoard.Snake), &snake)
				err = json.Unmarshal([]byte(selectBoard.Ladder), &ladder)
			}
			if getBoard != nil {
				snake = getBoard.Snake
				ladder = getBoard.Ladder
			}
			point := CalculateNextPoint(ctx, snake, ladder, &vo.Chess{
				Point: begin[len(begin)-1],
				Move:  throw[len(throw)-1],
			})
			err = redis.SetDiceByNextPonit(ctx, board, player, throw, begin, point)
			if err != nil {
				return 0, err
			}
		}
	} else if err != nil {
		return 0, err
	}
	randomInt := vo.NewDice()
	err = redis.SetDice(ctx, board, player, randomInt, 0, 0, 0, false)
	if err != nil {
		return 0, err
	}
	return randomInt, nil
}
