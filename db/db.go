package db

import (
	"log"

	"github.com/jinzhu/gorm"
	// _ - db connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
	m "github.com/samrocksc/gofunzi/models"
)

var db *gorm.DB
var err error

//Error Init - Auto Migrates up
func Init() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=happening sslmode=disable")
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	db.AutoMigrate(&m.User{}, m.Color{}, m.Event{})

	p1 := m.User{FirstName: "John", LastName: "Smith", Email: "testuser1@testdom.com"}
	p2 := m.User{FirstName: "Jane", LastName: "Smith", Email: "testuser2@testdom.com"}

	red := m.Color{Name: "Red", Description: "It's Red", Hex: "FF0000"}
	blue := m.Color{Name: "Blue", Description: "It's Blue", Hex: "0000FF"}
	green := m.Color{Name: "Green", Description: "It's Green", Hex: "00FF00"}

	db.Create(&p1)
	db.Create(&p2)
	db.Create(&red)
	db.Create(&blue)
	db.Create(&green)

}

//GetDB - return an instance of the DB
func GetDB() *gorm.DB {
	return db
}

// CloseDB - close current instance of the db
func CloseDB() {
	db.Close()
}
