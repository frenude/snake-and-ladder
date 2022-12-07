package service

import (
	"context"
	"snake-and-ladder/model/vo"
	"sort"
)

func CalculateNextPoint(ctx context.Context, snake, ladder [][]int, chess *vo.Chess) int {
	s := chess.Point + chess.Move
	if s >= 100 {
		return 100
	}
	for i := 0; i < len(snake); i++ {
		inSlice := snake[i]
		sort.Sort(sort.Reverse(sort.IntSlice(inSlice)))
		if s == inSlice[0] {
			return inSlice[1]
		}
	}
	for i := 0; i < len(ladder); i++ {
		inSlice := ladder[i]
		sort.Ints(inSlice)
		if s == inSlice[0] {
			return inSlice[1]
		}
	}
	return s
}
