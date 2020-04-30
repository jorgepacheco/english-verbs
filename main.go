package main

import (
	"encoding/json"
	"english-verbs/internal/verbs"
	"english-verbs/internal/verbs/persistence/csv"
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

	repository := csv.NewRepository()

	var s Specification
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := s.Port

	if port == "" {
		port = "80"
	}

	log.Println(":" + port)
	log.Println("😄 Init english3 verbs API ")
	http.ListenAndServe(":"+port, handler{repo: repository})
}
