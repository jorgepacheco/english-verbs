package main

import (
	"encoding/json"
	"english-verbs/internal/verbs"
	"english-verbs/internal/verbs/persistence/memory"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type handler struct {
	repo verbs.VerbRepo
}

func (h handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/verbs" && req.Method == http.MethodGet {
		log.Printf("Acceso a Verbos ...🏀")

		allVerbs, _ := h.repo.GetAllVerbs()
		writeBody(writer, allVerbs)
		writer.WriteHeader(http.StatusOK)
		return
	}

	log.Printf("No se pudo procesar el request ...")
	writer.WriteHeader(http.StatusNotFound)
	return
}

func writeBody(writer http.ResponseWriter, verbs []verbs.Verb) {
	/*	myBody := map[string]string{
			"verbs": "to be",
		}
	*/
	json.NewEncoder(writer).Encode(verbs)
}

type Specification struct {
	Debug      bool
	Port       string
	User       string
	Users      []string
	Rate       float32
	ColorCodes map[string]int
}

func main() {

	repository := memory.NewRepository()

	var s Specification
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(":" + s.Port)
	log.Println("😄 Init english2 verbs API ")
	http.ListenAndServe(":"+s.Port, handler{repo: repository})
}
