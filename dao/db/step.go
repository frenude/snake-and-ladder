package db

import (
	"context"
	"errors"
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
	result := dbClient.Where("bid=? AND u_name=?", boardId, player).Find(&findSteps)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	if findSteps == nil {
		return nil, nil, errors.New("查询数据不存在")
	}
	for _, i := range findSteps {
		begin = append(begin, int(i.Point))
		throw = append(throw, int(i.Next))
	}
	return begin, throw, nil
}
