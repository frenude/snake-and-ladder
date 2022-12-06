package vo

import (
	"fmt"
	"testing"
)

func TestGenSnakeBoard(t *testing.T) {
	m, n := 10, 10
	snake := GenSnakeBoard(m, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j != m-1 {
				fmt.Printf("%d ", snake[i][j])
			} else {
				fmt.Println(snake[i][j])
			}
		}
	}
}
func TestRandomPoint(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, RandomPoint(s))
	}
	t.Logf("%v", s)
}

func TestRandomSnakeAndLadder(t *testing.T) {
	snake, ladder := RandomSnakeAndLadder()
	t.Logf("%v", snake)
	t.Logf("%v", ladder)
}

func TestNewBoard(t *testing.T) {
	board, err := NewBoard(4)
	if err != nil {
		t.Logf("err:%v", err)
	}
	t.Logf("%v", board)
}
