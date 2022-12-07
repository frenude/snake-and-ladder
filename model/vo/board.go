package vo

import (
	"math/rand"
	"snake-and-ladder/utils"
	"time"
)

type BoardReq struct {
	PlayerNum int `json:"player_nums""`
}

type Board struct {
	Id       string
	Position [][]int
	Snake    [][]int
	Ladder   [][]int
	Player   []string
}

func NewBoard(n int, name string, snake, ladder [][]int) (*Board, error) {
	var player []string
	for i := 0; i < n; i++ {
		id := utils.GetId()
		token, err := utils.GenerateToken(name, id)
		if err != nil {
			return nil, err
		}
		player = append(player, token)
	}
	return &Board{
		Id:       name,
		Position: GenSnakeBoard(10, 10),
		Snake:    snake,
		Ladder:   ladder,
		Player:   player,
	}, nil
}

func RandomSnakeAndLadder() (snake, ladder [][]int) {
	rand.Seed(time.Now().UnixNano())
	// 随机出蛇的个数
	snakeNum := rand.Intn(5) + 5
	ladderNum := rand.Intn(5) + 5
	// 随机出蛇的位置 禁止重复
	var PointList []int
	for i := 0; i < 2*snakeNum+2*ladderNum; i++ {
		PointList = append(PointList, RandomPoint(PointList))
	}
	var row []int
	for i := 0; i < len(PointList); i++ {

		row = append(row, PointList[i])
		if len(row) == 2 {
			if len(snake) <= snakeNum {
				snake = append(snake, row)
			} else {
				ladder = append(ladder, row)
			}
			row = []int{}
		}

	}
	return snake, ladder
}
func RandomPoint(snakelist []int) int {
	rand.Seed(time.Now().UnixNano())
	snakePoint := rand.Intn(100)
	for _, s := range snakelist {
		if snakePoint == s {
			snakePoint = RandomPoint(snakelist)
		} else {
			return snakePoint
		}
	}
	return snakePoint
}
