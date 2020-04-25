package main

import (
	"encoding/json"
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

func main() {
	log.Println("😄 Init english verbs API ")
	http.ListenAndServe(":80", handler{})
}
