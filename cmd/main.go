package main

import (
	"english-verbs/api"
	"english-verbs/internal/logs"
	"fmt"
	"os"
)

const defaultPort = "80"

func main() {

	fmt.Println("::: Initial VERBS APP ")

	applicationPort := os.Getenv("PORT")

	logs.InitDefault("dev") //hardcoded at dev environment for the PoC API

	if applicationPort == "" {
		applicationPort = defaultPort
	}
	logs.Log().Info(applicationPort)

	api.Start(applicationPort)

}

/*func main()  {
	ctx := context.Background()

	repo := persistence.VerbRepository{}

	servicio := app.VerbService{Repo: &repo}

	all, _ := servicio.GetAll(ctx)

	fmt.Printf("%v", all)
}*/
