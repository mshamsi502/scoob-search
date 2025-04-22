package main

import (
	"fmt"

	"github.com/mshamsi502/scoob-search/fa_ar_characters"
	"github.com/mshamsi502/scoob-search/scoob"
)

func main() {
	// تست تبدیل کاراکترهای عربی به فارسی
	text := "أريد كاف و ياء و واء في النص"
	convertedText := fa_ar_characters.ArabicToPersianCharacter(text)
	fmt.Println("Converted Text:", convertedText)

	// استفاده از پکیج scoob برای جستجو
	scoobInstance := scoob.Scoob{}
	objects := []string{"شهر", "کتاب", "آسمان"}
	query := "کتاب"
	result := scoobInstance.DoobySearch(query, objects)
	fmt.Println("Search Results:", result)
}
