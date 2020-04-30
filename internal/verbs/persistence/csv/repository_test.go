package csv

import "testing"

func TestRepository_GetAllVerbs(t *testing.T) {

	expected := 1

	repo := NewRepository()

	verbs, _ := repo.GetAllVerbs()

	if len(verbs) != expected {
		t.Errorf("Expected: %v, got: %v", expected, len(verbs))
	}
}
