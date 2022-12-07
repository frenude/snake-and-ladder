package service

import (
	"context"
	"encoding/json"
	"errors"
	goredis "github.com/go-redis/redis/v8"
	"snake-and-ladder/dao/db"
	"snake-and-ladder/dao/redis"
	"snake-and-ladder/model/po"
	"snake-and-ladder/model/vo"
	"strconv"
)

func Step(ctx context.Context, board, player string, chess *vo.Chess) (int, error) {
	// 判断是否作弊
	dice, err := redis.GetDice(ctx, board, player)
	if err == goredis.Nil || dice.Next != chess.Move || dice.NextPoint != chess.Point {
		return 0, errors.New("存在作弊嫌疑，请重新按照标准输入")
	}
	if dice.NextPoint == 100 {
		return 0, errors.New("游戏结束，请勿投掷骰子")
	}
	atoi, err := strconv.Atoi(board)
	if err != nil {
		return 0, err
	}
	err = db.CreateStep(ctx, &po.Step{
		Point: uint8(chess.Point),
		Next:  uint8(chess.Move),
		Bid:   uint(atoi),
		UName: player,
	})
	if err != nil {
		return 0, err
	}
	// 查询棋盘结构
	var snake [][]int
	var ladder [][]int
	getBoard, err := redis.GetBoard(ctx, board)
	if err == goredis.Nil {
		//redis 查不到mysql查询
		i, err := strconv.Atoi(board)
		if err != nil {
			return 0, err
		}
		selectBoard, err := db.SelectBoard(ctx, uint(i))
		if err != nil {
			return 0, err
		}
		err = json.Unmarshal([]byte(selectBoard.Snake), &snake)
		err = json.Unmarshal([]byte(selectBoard.Ladder), &ladder)

		err = redis.SetBoard(ctx, &vo.Board{
			Id:       strconv.Itoa(int(selectBoard.Id)),
			Position: nil,
			Snake:    snake,
			Ladder:   ladder,
			Player:   []string{player},
		})
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}
	if getBoard != nil {
		snake = getBoard.Snake
		ladder = getBoard.Ladder
	}
	// 计算位置
	next_point := CalculateNextPoint(ctx, snake, ladder, chess)
	// 存储nextPoint

	if dice.NextPoint == next_point {
		return 0, errors.New("重复输入请重新投掷骰子")
	}
	err = redis.SetDice(ctx, board, player, chess.Move, chess.Move, chess.Point, next_point, true)
	if err != nil {
		return 0, err
	}
	return next_point, nil
}
