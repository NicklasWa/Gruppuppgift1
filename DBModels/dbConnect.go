package DBModels

import (
	"github.com/glebarez/sqlite"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func Open_SQLite() {

	db, err := gorm.Open(sqlite.Open("SQLite_files/team_players.sqlite"), &gorm.Config{})
	if err != nil {
		panic("faild to connect database")
	}

	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Team{})
}
