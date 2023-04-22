package DBModels

type Player struct {
	ID          uint   `gorm:"primaryKey"`
	FoundedYear uint32 `gorm:"type:integer"`
	BirthYear   uint32 `gorm:"type:integer"`
	Name        string `gorm:"type:text"`
	City        string `gorm:"type:text"`
}
