package vo

type Game struct {
	Id       string
	Position [][]int
	Snake    [][]int
	Ladder   [][]int
	Throw    []int
	Begin    []int
}
