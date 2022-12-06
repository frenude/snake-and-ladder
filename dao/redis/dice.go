package redis

import (
	"context"
	"encoding/json"
	"snake-and-ladder/model/dto"
	"strings"
	"time"
)

func SetDice(ctx context.Context, board, player string, next, throw, begin int) error {
	name := strings.Builder{}
	name.WriteString(board)
	name.WriteString(":")
	name.WriteString(player)
	rediskey := name.String()
	result, err := redisClient.Get(ctx, rediskey).Result()
	var play *dto.Player
	if err != nil {
		play = &dto.Player{
			Next:  next,
			Throw: []int{},
			Begin: []int{},
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
	var p dto.Player
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		return err
	}
	if next != 0 {
		if begin == 0 && throw == 0 {
			play = &dto.Player{
				Next:  next,
				Throw: []int{},
				Begin: []int{},
			}
		} else if throw != 0 {
			t := append(p.Throw, throw)
			play = &dto.Player{
				Next:  next,
				Throw: t,
				Begin: p.Begin,
			}
		} else if begin != 0 || (begin == 0 && len(p.Begin) == 0) {
			b := append(p.Begin, begin)
			play = &dto.Player{
				Next:  next,
				Throw: p.Throw,
				Begin: b,
			}
		} else {
			b := append(p.Begin, begin)
			t := append(p.Throw, throw)
			play = &dto.Player{
				Next:  next,
				Throw: t,
				Begin: b,
			}
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
