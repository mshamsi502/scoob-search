package scoob

import "fmt"

// ساختار Scoob برای جستجو
type Scoob struct{}

// جستجو در لیست با توجه به رشته وارد شده
func (s Scoob) DoobySearch(query string, objects []string) []string {
	var result []string
	for _, obj := range objects {
		if containsIgnoreCase(obj, query) {
			result = append(result, obj)
		}
	}
	return result
}

// جستجو با غفلت از حروف بزرگ و کوچک
func containsIgnoreCase(text, substr string) bool {
	return (len(text) >= len(substr)) && (stringMatch(text, substr))
}

// تطابق رشته‌ها
func stringMatch(text, substr string) bool {
	return stringIndex(text, substr) != -1
}

// پیدا کردن ایندکس زیررشته در رشته اصلی
func stringIndex(text, substr string) int {
	for i := 0; i <= len(text)-len(substr); i++ {
		if text[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// تست اولیه برای چک کردن وضعیت پکیج
func Test() {
	fmt.Println("Scoob Package Ready 🐾")
}
