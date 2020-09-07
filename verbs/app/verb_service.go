package app

import (
	"context"
	"english-verbs/verbs/domain"
	"fmt"
	"math/rand"
	"time"
)

// Interfaz

type VerbService struct {
	Repo domain.Repository
}

func (srv *VerbService) GetAll(ctx context.Context) ([]domain.Verb, error) {

	return srv.Repo.GetAll(ctx)
}

func (srv *VerbService) GetExam(numberOfQuestions int) ([]domain.Verb, error) {

	var exam []domain.Verb

	questions := generateQuestions(numberOfQuestions, srv.Repo.GetSize())

	fmt.Printf(":: Questions %v", questions)

	for _, question := range questions {

		examElem, _ := srv.Repo.FindById(question)

		exam = append(exam, examElem)
	}

	return exam, nil
}

func generateQuestions(numberOfQuestions, limit int) []int {

	var questions []int

	for len(questions) < numberOfQuestions {
		questions = append(questions, random(limit))
	}
	return questions
}

func random(limit int) int {

	seed := rand.NewSource(time.Now().UnixNano())

	return rand.New(seed).Intn(limit)

}
