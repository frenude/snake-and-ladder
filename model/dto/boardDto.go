package dto

type Player struct {
	Next  int   `json:"next"`
	Throw []int `json:"throw"`
	Begin []int `json:"begin"`
}

type BoardDTO struct {
	Snake   [][]int  `json:"snake"`
	Ladder  [][]int  `json:"ladder"`
	Players []string `json:"players"`
}
