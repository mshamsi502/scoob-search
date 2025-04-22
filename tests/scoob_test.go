package tests

import (
	"testing"

	"github.com/mshamsi502/scoob-search/scoob"
)

func TestDoobySearch(t *testing.T) {
	s := scoob.Scoob{}
	list := []string{"سعید", "احمد", "حمید", "امیرحسین", "فاطمه"}

	result := s.DoobySearch("م", list)
	expectedCount := 4

	if len(result) != expectedCount {
		t.Errorf("Expected %d results, but got %d", expectedCount, len(result))
	}
}
