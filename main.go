package main

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type handler struct {
}

func (h handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/verbs" && req.Method == http.MethodGet {
		writeBody(writer)
		writer.WriteHeader(http.StatusOK)
		return
	}

	log.Printf("No se pudo procesar el request")
	writer.WriteHeader(http.StatusNotFound)
	return
}

func writeBody(writer http.ResponseWriter) {
	myBody := map[string]string{
		"hello": "gophers",
	}

	json.NewEncoder(writer).Encode(myBody)
}


type Specification struct {
	Debug       bool
	Port        string
	User        string
	Users       []string
	Rate        float32
	ColorCodes  map[string]int
}

func main() {


	var s Specification
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(":" + s.Port)
	log.Println("😄 Init english2 verbs API ")
	http.ListenAndServe(":" + s.Port, handler{})
}
