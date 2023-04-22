package DBModels

import (
	"github.com/glebarez/sqlite"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB // Publik variabel som håller en pekare till den skapade SQLite-databasen

func Open_SQLite() {

	db, err := gorm.Open(sqlite.Open("SQLite_files/team_players.sqlite"), &gorm.Config{})
	if err != nil {
		panic("faild to connect database")
	}

	DB = db

	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Team{})

	// Seed team data
	db.Create(&Team{FoundedYear: 1878, Name: "Manchester United", City: "Manchester"})
	db.Create(&Team{FoundedYear: 1892, Name: "Liverpool", City: "Liverpool"})
	db.Create(&Team{FoundedYear: 1886, Name: "Arsenal", City: "London"})
	db.Create(&Team{FoundedYear: 1905, Name: "Chelsea", City: "London"})
	db.Create(&Team{FoundedYear: 1882, Name: "Tottenham Hotspur", City: "London"})
	db.Create(&Team{FoundedYear: 1880, Name: "Manchester City", City: "Manchester"})
	db.Create(&Team{FoundedYear: 1878, Name: "Everton", City: "Liverpool"})
	db.Create(&Team{FoundedYear: 1874, Name: "Aston Villa", City: "Birmingham"})
	db.Create(&Team{FoundedYear: 1892, Name: "Newcastle United", City: "Newcastle upon Tyne"})
	db.Create(&Team{FoundedYear: 1919, Name: "Leeds United", City: "Leeds"})
}
