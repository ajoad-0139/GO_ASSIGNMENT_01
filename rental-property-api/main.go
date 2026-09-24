package main

import (
	"rental-property-api/services"
	"time"

	_ "rental-property-api/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// constraints for retry
const (
	maxRetries = 5
	retryDelay = 3 * time.Second
)

func main() {

	// load data before starting the app
	err := services.LoadDataWithRetry(maxRetries, retryDelay)
	if err != nil {
		logs.Critical("Error : Failed to load data %v", err)
		panic(err)
	}

	// serve the generated swagger.json/yaml + UI assets from the swagger/ folder
	if beego.BConfig.RunMode == "dev" {
		beego.SetStaticPath("/swagger", "swagger")
	}

	//start server
	beego.Run()
}
