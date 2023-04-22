package DBModels

type Team struct {
	ID          uint   `gorm:"primaryKey"`
	FoundedYear int    `gorm:"type:integer"`
	Name        string `gorm:"type:text"`
	City        string `gorm:"type:text"`
}
