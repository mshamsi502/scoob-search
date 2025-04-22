package scoob

import "fmt"

type Scoob struct{}

func (s Scoob) DoobySearch(query string, objects []string) []string {
	var result []string
	for _, obj := range objects {
		if containsIgnoreCase(obj, query) {
			result = append(result, obj)
		}
	}
	return result
}

func containsIgnoreCase(text, substr string) bool {
	return (len(text) >= len(substr)) && (stringMatch(text, substr))
}

func stringMatch(text, substr string) bool {
	return stringIndex(text, substr) != -1
}

func stringIndex(text, substr string) int {
	for i := 0; i <= len(text)-len(substr); i++ {
		if text[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func Test() {
	fmt.Println("Scoob Package Ready 🐾")
}
