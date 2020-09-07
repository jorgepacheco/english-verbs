package main

import (
	"english-verbs/api"
	"english-verbs/internal/logs"
	"fmt"
	"math/rand"
	"os"
	"time"
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

func main4() {

	for i := 0; i < 10; i++ {
		//rand.Seed(time.Now().UnixNano())
		//fmt.Printf("%v,", rand.Intn(100))
		x1 := rand.NewSource(time.Now().UnixNano())

		y1 := rand.New(x1)

		fmt.Println(y1.Intn(360))

	}

}
