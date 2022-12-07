package redis

import (
	"context"
	"encoding/json"
	"snake-and-ladder/model/dto"
	"snake-and-ladder/model/vo"
	"time"
)

func SetBoard(ctx context.Context, board *vo.Board) error {
	redisBoard := &dto.BoardDTO{
		Snake:   board.Snake,
		Ladder:  board.Ladder,
		Players: board.Player,
	}
	marshal, err := json.Marshal(redisBoard)
	if err != nil {
		return err
	}
	err = redisClient.Set(ctx, board.Id, marshal, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetBoard(ctx context.Context, name string) (*dto.BoardDTO, error) {
	result, err := redisClient.Get(ctx, name).Result()
	if err != nil {
		return nil, err
	}
	var board dto.BoardDTO
	err = json.Unmarshal([]byte(result), &board)
	if err != nil {
		return nil, err
	}
	return &board, nil
}
