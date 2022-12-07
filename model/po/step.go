package po

type Step struct {
	Id    uint `gorm:"primaryKey"`
	Point uint8
	Next  uint8
	Bid   uint
	UName string
}
