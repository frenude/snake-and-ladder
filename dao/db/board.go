package db

import (
	"context"
	"snake-and-ladder/model/po"
)

func CreateBoard(ctx context.Context, board *po.Board) error {
	result := dbClient.Create(board)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func SelectBoards(ctx context.Context) ([]*po.Board, error) {
	var findBorads []*po.Board
	result := dbClient.Find(&findBorads)
	if result.Error != nil {
		return nil, result.Error
	}
	return findBorads, nil
}
func SelectBoard(ctx context.Context, id uint) (*po.Board, error) {
	var findBorad *po.Board
	result := dbClient.First(&findBorad, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return findBorad, nil
}
