package po

type Board struct {
	Id     uint `gorm:"primaryKey"`
	Snake  string
	Ladder string
	TimeInfo
}
