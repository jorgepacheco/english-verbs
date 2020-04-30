package memory

import (
	"english-verbs/internal/verbs"
	"fmt"
)

var verbsStorage = map[string]verbs.Verb{
	"START": {Infinitive: "START", Past: "STARTED", Participle: "STARTED", Meaning: "Comenzar - Arrancar"},
	"STEAL": {Infinitive: "STEAL", Past: "STOLE", Participle: "STOLEN", Meaning: "Robar - Hurtar"},
	"STINK": {Infinitive: "START", Past: "STANK", Participle: "STUNK", Meaning: "Apestar - Oler mal"},
}

func init() {
	fmt.Println(":: Init Storage in Memory")
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
