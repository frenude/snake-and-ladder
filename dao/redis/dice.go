package redis

import (
	"context"
	"encoding/json"
	"errors"
	goredis "github.com/go-redis/redis/v8"
	"snake-and-ladder/model/dto"
	"strings"
	"time"
)

func SetDice(ctx context.Context, board, player string, next, throw, begin, next_point int, isStep bool) error {
	name := strings.Builder{}
	name.WriteString(board)
	name.WriteString(":")
	name.WriteString(player)
	rediskey := name.String()
	result, err := redisClient.Get(ctx, rediskey).Result()
	var play *dto.Player
	if err == goredis.Nil {
		if err != nil {
			play = &dto.Player{
				Next:      next,
				Throw:     []int{},
				Begin:     []int{},
				NextPoint: next_point,
			}
			marshal, err := json.Marshal(play)
			if err != nil {
				return err
			}
			err = redisClient.Set(ctx, rediskey, marshal, time.Hour).Err()
			if err != nil {
				return err
			}
			return nil
		} else if err != nil {
			return err
		}
	}
	var p dto.Player
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		return err
	}
	if p.NextPoint == 100 {
		return errors.New("游戏结束，请勿投掷骰子")
	}
	if !isStep {
		play = &dto.Player{
			Next:      next,
			Throw:     p.Throw,
			Begin:     p.Begin,
			NextPoint: p.NextPoint,
		}
	} else {
		b := append(p.Begin, begin)
		t := append(p.Throw, throw)
		play = &dto.Player{
			Next:      next,
			Throw:     t,
			Begin:     b,
			NextPoint: next_point,
		}
	}

	marshal, err := json.Marshal(play)
	if err != nil {
		return err
	}
	err = redisClient.Set(ctx, rediskey, marshal, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetDice(ctx context.Context, board, player string) (*dto.Player, error) {
	name := strings.Builder{}
	name.WriteString(board)
	name.WriteString(":")
	name.WriteString(player)
	rediskey := name.String()
	result, err := redisClient.Get(ctx, rediskey).Result()
	if err != nil {
		return nil, err
	}
	var p dto.Player
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
