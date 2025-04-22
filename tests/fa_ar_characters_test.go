package tests

import (
	"testing"

	"github.com/mshamsi502/scoob-search/fa_ar_characters"
)

func TestArabicToPersianCharacter(t *testing.T) {
	input := "أريد كاف و ياء و واء"
	expected := "ارید کاف و یاء و واء"

	result := fa_ar_characters.ArabicToPersianCharacter(input)
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
