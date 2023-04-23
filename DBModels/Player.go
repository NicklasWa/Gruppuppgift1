package DBModels

type Player struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	TeamID       uint   `json:"team_id" gorm:"index"`
	Team         *Team  `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	JerseyNumber int    `json:"jersey_number" gorm:"type:integer"`
	Name         string `json:"name" gorm:"type:text"`
	City         string `json:"city" gorm:"type:text"`
	BirthYear    int    `json:"birth_year" gorm:"type:integer"`
}
