package domain

type Verb struct {
	Infinitive string `json:"infinitive"`
	Past       string `json:"past"`
	Participle string `json:"participle"`
	Meaning    string `json:"meaning"`
}

func NewVerb(infinitive, past, participle, meaning string) (verb Verb) {

	verb = Verb{
		Infinitive: infinitive,
		Past:       past,
		Participle: participle,
		Meaning:    meaning,
	}
	return
}
