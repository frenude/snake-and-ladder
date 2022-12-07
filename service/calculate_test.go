package service

import (
	"context"
	"snake-and-ladder/model/vo"
	"testing"
)

func TestCalculateNextPoint(t *testing.T) {
	snake := [][]int{{22, 96}, {42, 29}, {40, 0}, {8, 95}, {17, 65}, {40, 95}}
	ladder := [][]int{{63, 11}, {14, 44}, {82, 61}, {3, 78}, {25, 78}, {61, 53}, {95, 29}}
	chess := vo.Chess{
		Point: 63,
		Move:  2,
	}
	var ctx context.Context
	point := CalculateNextPoint(ctx, snake, ladder, &chess)
	t.Logf("point : %v", point)
}
