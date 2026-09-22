package main

import (
	"rental-property-api/services"

	_ "rental-property-api/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	// load data before starting the app
	err := services.LoadData()
	if err != nil {
		logs.Critical("Failed to load data %v", err)
		panic(err)
	}

	beego.Run()
}
