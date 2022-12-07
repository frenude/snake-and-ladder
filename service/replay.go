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

func Replay(ctx context.Context, board string) (*vo.Game, error) {
	var game vo.Game
	var players []string
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
		users, err := db.SelectUsers(ctx, uint(atoi))
		players = users
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
		players = getBoard.Players
	}
	var ps []vo.Player
	for i := 0; i < len(players); i++ {
		var s vo.Player
		player := players[i]
		s.Name = player
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
			s.Throw = throw
			s.Begin = begin
		} else if err != nil {
			return nil, err
		}
		if dice != nil {
			s.Begin = dice.Begin
			s.Throw = dice.Throw
		}
		ps = append(ps, s)
	}
	game.Player = ps
	return &game, nil
}
