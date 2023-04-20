package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"systementor.se/yagolangapi/data"
)

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "test", Rubrik: "Hej Golang"})
}

var theRandom *rand.Rand

// HTML
// JSON

func mekashaJson(c *gin.Context) {
	var listofPeople = []data.Person{
		{Name: "Mekasha", City: "Stockholm"},
	}
	c.IndentedJSON(http.StatusOK, listofPeople)
}

func nicklasJson(c *gin.Context) {
	var nicklas data.Person
	nicklas.Name = "Nicklas"
	nicklas.City = "Falun"

	c.JSON(http.StatusOK, nicklas)
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
