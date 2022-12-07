package assembler

import (
	"web/model/dto"
	"web/model/po"
)

/**
 * @author: Frenude
 * @mail: frenude@gmail.com
 * @date: created in 04 23,2022
 * @desc: //todo
 */
func UsersToPo(old []*dto.User) []*po.User {
	result := make([]*po.User, len(old))
	for i, v := range old {
		result[i] = &po.User{Name: v.UserName, Age: v.UserAge}
	}
	return result
}

// UserToPo *pb.User 转换为 *po.User
func UserToPo(old *dto.User) *po.User {
	return &po.User{
		Id:   old.UserId,
		Name: old.UserName,
		Age:  old.UserAge,
	}
}
