package csv

import (
	"bufio"
	"encoding/csv"
	"english-verbs/internal/verbs"
	"fmt"
	"io"
	"log"
	"os"
)

var verbsStorage = map[string]verbs.Verb{}

func init() {
	fmt.Println(":: Init CSV Storage")
	data := readData()
	for _, verb := range data {
		verbsStorage[verb.Infinitive] = verb
	}

	//verb := verbs.NewVerb("START", "STARTED", "STARTED", "Comenzar - Arrancar")
	//verbsStorage[verb.Infinitive] = verb
}

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() verbs.VerbRepo {
	return &repository{}
}

func (repo repository) GetAllVerbs() ([]verbs.Verb, error) {
	var verbs []verbs.Verb
	for _, value := range verbsStorage {
		verbs = append(verbs, value)
	}
	return verbs, nil
}

func readData() []verbs.Verb {
	csvFile, error := os.Open("./data/all_verbs.csv")

	if error != nil {
		log.Println("Error read csv Error", error)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	var verbsItems []verbs.Verb
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		verbsItems = append(verbsItems, verbs.Verb{
			Infinitive: line[1],
			Past:       line[2],
			Participle: line[3],
			Meaning:    line[4],
		})
	}

	return verbsItems
}
