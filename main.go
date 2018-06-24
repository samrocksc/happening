package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	db "github.com/samrocksc/gofunzi/db"
	routes "github.com/samrocksc/gofunzi/routes"
)

func main() {
	db.Init()
	r := routes.SetupRouter()
	r.Run(":8080")
}
