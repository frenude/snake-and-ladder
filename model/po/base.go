package po

import (
	"gorm.io/gorm"
	"time"
)

// TimeInfo 被嵌入到其他表结构中，为记录提供操作时间信息
// 遵循 gorm 的约定，使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间，gorm 会自动为这些字段赋值
// 软删除需使用 gorm.DeletedAt 类型，执行 Delete 方法时，gorm 会将 DeletedAt 置为当前时间
type TimeInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
