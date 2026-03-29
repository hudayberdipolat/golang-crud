package main

import (
	"fmt"
	"log"

	"github.com/hudayberdipolat/golang-crud/internal/app"
	"github.com/hudayberdipolat/golang-crud/internal/setup/constructor"
)

func main() {
	getDependencies, err := app.GetAppDependencies()
	if err != nil {
		log.Fatalf("error %v", err.Error())
		return
	}
	constructor.Build(getDependencies)
	server := fmt.Sprintf("%s:%s",
		getDependencies.Config.ServerConfig.ServerHost,
		getDependencies.Config.ServerConfig.ServerPort)
	newApp := app.NewApp(getDependencies)
	if errServerRun := newApp.Listen(server); errServerRun != nil {
		log.Fatal(errServerRun.Error())
		return
	}

}
