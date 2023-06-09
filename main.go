package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"systementor.se/yagolangapi/DBModels"
	"systementor.se/yagolangapi/data"
)

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "test", Rubrik: "Hej Golang"})
}

var theRandom *rand.Rand

// HTML
// JSON
// Get Mekasha
func mekashaJson(c *gin.Context) {
	var listofPeople = []data.Person{
		{Name: "Mekasha", City: "Stockholm"},
	}
	c.IndentedJSON(http.StatusOK, listofPeople)
}

// Get Nicklas
func nicklasJson(c *gin.Context) {
	var nicklas data.Person
	nicklas.Name = "Nicklas"
	nicklas.City = "Falun"

	c.JSON(http.StatusOK, nicklas)
}

// Get All teams
func teamsJson(c *gin.Context) {
	var teams []DBModels.Team // En tom lista av team skapas
	DBModels.DB.Find(&teams)  // Hämta data från databasen.

	c.IndentedJSON(http.StatusOK, teams)
}

// Get Team by ID
func getTeamByID(c *gin.Context) {
	var team DBModels.Team
	if err := DBModels.DB.Preload("Players").First(&team, c.Param("teamid")).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	var playerNames []string
	for _, player := range team.Players {
		playerNames = append(playerNames, player.Name)
	}

	response := gin.H{
		"Name":        team.Name,
		"City":        team.City,
		"FoundedYear": team.FoundedYear,
		"Players":     playerNames,
	}
	c.JSON(http.StatusOK, response)

}

// Get Player by ID
func getPlayerByID(c *gin.Context) {
	id := c.Param("playerid")
	var player DBModels.Player
	var team DBModels.Team
	var teamid uint

	DBModels.DB.Find(&player, "id= ?", id)

	teamid = player.TeamID
	DBModels.DB.Find(&team, "id= ?", teamid)
	teamName := team.Name

	response := gin.H{
		"Name":         player.Name,
		"Team":         teamName,
		"JerseyNumber": player.JerseyNumber,
		"BirthYear":    player.BirthYear,
	}

	c.IndentedJSON(http.StatusOK, response)

}

func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func addEmployee(c *gin.Context) {

	data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})

}

func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
	}

}

type PageView struct {
	Title  string
	Rubrik string
}

var config Config

func main() {

	theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	readConfig(&config)
	DBModels.Open_SQLite()
	data.InitDatabase(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/nicklas", nicklasJson)
	router.GET("/api/mekasha", mekashaJson)
	router.GET("/api/teams", teamsJson)
	router.GET("/api/team/:teamid", getTeamByID)
	router.GET("/api/player/:playerid", getPlayerByID)
	router.GET("/api/employees", employeesJson)
	router.GET("/api/addemployee", addEmployee)
	router.GET("/api/addmanyemployees", addManyEmployees)
	router.Run("localhost:8080")

	// e := data.Employee{
	// 	Age:  1,
	// 	City: "Strefabn",
	// 	Namn: "wddsa",
	// }

	// if e.IsCool() {
	// 	fmt.Printf("Namn is cool:%s\n", e.Namn)
	// } else {
	// 	fmt.Printf("Namn:%s\n", e.Namn)
	// }

	// fmt.Println("Hello")
	// t := tabby.New()
	// t.AddHeader("Namn", "Age", "City")
	// t.AddLine("Stefan", "50", "Stockholm")
	// t.AddLine("Oliver", "14", "Stockholm")
	// t.Print()
}
