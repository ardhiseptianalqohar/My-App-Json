package main

import (
	// "fmt"
	"myapp/config"
	"myapp/routes"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()

	// fmt.Println("====================================")
	// fmt.Println("\tSELAMAT DATANG BOSS..")
	// fmt.Println("===================================")

	e := routes.Init()
	e.Logger.Fatal(e.Start(":2109"))
}
