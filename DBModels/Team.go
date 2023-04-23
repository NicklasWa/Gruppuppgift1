package DBModels

type Team struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	FoundedYear int      `json:"founded_year" gorm:"type:integer"`
	Name        string   `json:"name" gorm:"type:text"`
	City        string   `json:"city" gorm:"type:text"`
	Players     []Player `json:"players,omitempty" gorm:"foreignKey:TeamID"`
}
