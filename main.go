package main

import (
	"fmt"

	"github.com/mshamsi502/scoob-search/fa_ar_characters"
	"github.com/mshamsi502/scoob-search/scoob"
)

func main() {
	// تست کانورت عربی به فارسی
	text := "أريد كاف و ياء و واء في النص"
	convertedText := fa_ar_characters.ArabicToPersianCharacter(text)
	fmt.Println("Converted Text: ", convertedText)

	// تست جستجو
	s := scoob.Scoob{}
	objects := []string{"سلام", "دنیا", "علیکم", "أريد", "کاف"}
	result := s.DoobySearch("كاف", objects)
	fmt.Println("Search Result: ", result)
	s2 := scoob.Scoob{}
	result2 := s2.DoobySearch("اريد", objects)
	fmt.Println("Search Result: ", result2)

	// تست وضعیت پکیج scoob
	scoob.Test()
}
