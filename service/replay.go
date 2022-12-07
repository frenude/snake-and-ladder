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

func Replay(ctx context.Context, board, player string) (*vo.Game, error) {
	var game vo.Game
	game.Id = board
	game.Position = vo.GenSnakeBoard(10, 10)
	getBoard, err := redis.GetBoard(ctx, board)
	if err == goredis.Nil {
		atoi, err := strconv.Atoi(board)
		if err != nil {
			return nil, err
		}
		selectBoard, err := db.SelectBoard(ctx, uint(atoi))
		if err != nil {
			return nil, err
		}
		var s [][]int
		err = json.Unmarshal([]byte(selectBoard.Snake), &s)
		if err != nil {
			return nil, err
		}
		var l [][]int
		err = json.Unmarshal([]byte(selectBoard.Ladder), &l)
		if err != nil {
			return nil, err
		}
		game.Snake = s
		game.Ladder = l
	} else if err != nil {
		return nil, err
	}
	if getBoard != nil {
		game.Snake = getBoard.Snake
		game.Ladder = getBoard.Ladder
	}
	dice, err := redis.GetDice(ctx, board, player)
	if err == goredis.Nil {
		atoi, err := strconv.Atoi(board)
		if err != nil {
			return nil, err
		}
		begin, throw, err := db.SelectStep(ctx, uint(atoi), player)
		if err != nil {
			return nil, err
		}
		game.Throw = throw
		game.Begin = begin
	} else if err != nil {
		return nil, err
	}
	if dice != nil {
		game.Begin = dice.Begin
		game.Throw = dice.Throw
	}
	return &game, nil
}
