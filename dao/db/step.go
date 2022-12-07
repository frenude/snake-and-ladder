package db

import (
	"context"
	"snake-and-ladder/model/po"
)

func CreateStep(ctx context.Context, step *po.Step) error {
	result := dbClient.Create(step)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SelectStep(ctx context.Context, boardId uint, player string) (begin []int, throw []int, err error) {
	var findSteps []*po.Step
	result := dbClient.Where("bid =", boardId).Where("uname =", player).Find(findSteps)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	for _, i := range findSteps {
		begin = append(begin, int(i.Point))
		throw = append(begin, int(i.Next))
	}
	return begin, throw, nil
}
