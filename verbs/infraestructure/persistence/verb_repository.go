package persistence

import (
	"bufio"
	"context"
	"encoding/csv"
	"english-verbs/verbs/domain"
	"fmt"
	"io"
	"log"
	"os"
)

var verbsStorage = map[string]domain.Verb{}

func init() {
	fmt.Println(":: Init CSV Storage")
	data := readData()
	for _, verb := range data {
		verbsStorage[verb.Infinitive] = verb
	}
	//verb := verbs.NewVerb("START", "STARTED", "STARTED", "Comenzar - Arrancar")
	//verbsStorage[verb.Infinitive] = verb
}

type VerbRepository struct {
}

func (repo *VerbRepository) FindById(id int) (domain.Verb, error) {

	all, _ := repo.GetAll(context.Background())

	return all[id], nil
}

func (repo *VerbRepository) GetAll(ctx context.Context) ([]domain.Verb, error) {

	var verbs []domain.Verb
	for _, value := range verbsStorage {
		verbs = append(verbs, value)
	}
	return verbs, nil
}

func readData() []domain.Verb {
	csvFile, error := os.Open("./data/all_verbs.csv")

	if error != nil {
		log.Println("Error read csv Error", error)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	var verbsItems []domain.Verb
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		verbsItems = append(verbsItems, domain.Verb{
			Infinitive: line[1],
			Past:       line[2],
			Participle: line[3],
			Meaning:    line[4],
		})
	}

	return verbsItems
}
