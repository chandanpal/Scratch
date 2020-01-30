package main

import (
	//"github.com/cisco-runner/config"
	//"github.com/cisco-runner/logger"
	"github.com/discover_services/app"
	//"log"
)


func main() {	
	app := &app.App{}
	app.Initialize()
	app.Run()
}
