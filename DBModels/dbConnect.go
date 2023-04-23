package DBModels

import (
	"github.com/glebarez/sqlite"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB // Publik variabel som håller en pekare till den skapade SQLite-databasen

func Open_SQLite() {

	db, err := gorm.Open(sqlite.Open("SQLite_files/Team_and_Players.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	var count int64
	db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table'").Count(&count)

	if count == 0 {
		seedData()
		return
	}
	return
}

func seedData() {
	DB.AutoMigrate(&Team{}, &Player{})

	// Seed team data
	team1 := Team{FoundedYear: 1878, Name: "Manchester United", City: "Manchester"}
	team2 := Team{FoundedYear: 1892, Name: "Liverpool", City: "Liverpool"}
	team3 := Team{FoundedYear: 1886, Name: "Arsenal", City: "London"}
	team4 := Team{FoundedYear: 1905, Name: "Chelsea", City: "London"}

	DB.Create(&team1)
	DB.Create(&team2)
	DB.Create(&team3)
	DB.Create(&team4)
	// Seed player data
	player1 := Player{TeamID: team1.ID, JerseyNumber: 25, Name: "Jaden Sancho", City: "London", BirthYear: 2000}
	player2 := Player{TeamID: team4.ID, JerseyNumber: 19, Name: "Mason Mount", City: "Portsmouth", BirthYear: 1999}
	player3 := Player{TeamID: team2.ID, JerseyNumber: 10, Name: "Mohamed Salah", City: "Nagrig", BirthYear: 1992}
	player4 := Player{TeamID: team3.ID, JerseyNumber: 8, Name: "Martin Odegaard", City: "Drammen", BirthYear: 1998}
	player5 := Player{TeamID: team1.ID, JerseyNumber: 10, Name: "Marcus Rashford", City: "Manchester", BirthYear: 1997}
	player6 := Player{TeamID: team2.ID, JerseyNumber: 21, Name: "Darwin Nunez", City: "Artigas", BirthYear: 1999}
	player7 := Player{TeamID: team4.ID, JerseyNumber: 7, Name: "N'Golo Kante", City: "Paris", BirthYear: 1991}
	player8 := Player{TeamID: team3.ID, JerseyNumber: 7, Name: "Bakayo Saka", City: "London", BirthYear: 2001}

	DB.Create(&player1)
	DB.Create(&player2)
	DB.Create(&player3)
	DB.Create(&player4)
	DB.Create(&player5)
	DB.Create(&player6)
	DB.Create(&player7)
	DB.Create(&player8)
}
