package main

import (
	"final-project/config"
	"final-project/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}
