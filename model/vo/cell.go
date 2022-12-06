package vo

type Cell struct {
	x int
	y int
}

func (c *Cell) Rigth() {
	c.x++
}

func (c *Cell) Down() {
	c.y++
}

func (c *Cell) Left() {
	c.x--
}
func (c *Cell) Up() {
	c.y--
}

func GenSnakeBoard(m, n int) [][]int {
	var snake [][]int
	cell := new(Cell)
	//初始化二维数组
	for i := 0; i < n; i++ {
		row := make([]int, m)
		snake = append(snake, row)
	}
	flag := 0 //flag表示当前走向，

	for i := m * n; i > 0; i-- { //填数据
		snake[cell.y][cell.x] = i //注意这里y代表的是列，x代表的是行
		switch flag {
		case 0:
			if (cell.x+1 >= m) || (snake[cell.y][cell.x+1] != 0) {
				flag = 1
				cell.Down()
			} else {
				cell.Rigth()
			}
		case 1:
			if (cell.x-1 < 0) || (snake[cell.y][cell.x-1] != 0) {
				flag = 0
				cell.Down()
			} else {
				cell.Left()
			}

		}
	}
	return snake
}
