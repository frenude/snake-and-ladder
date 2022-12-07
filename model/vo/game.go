package vo

type Player struct {
	Name  string
	Throw []int
	Begin []int
}

type Game struct {
	Id       string
	Position [][]int
	Snake    [][]int
	Ladder   [][]int
	Player   []Player
}
