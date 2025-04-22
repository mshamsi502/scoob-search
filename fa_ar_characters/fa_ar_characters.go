package scoob

import (
	"fmt"
	"strings"
)

// تعریف مقادیر کاراکترهای فارسی و عربی به صورت بایت
var ALEF_FA = []byte{216, 167}    // ا
var ALEF_AR_01 = []byte{216, 162} // آ
var ALEF_AR_02 = []byte{216, 163} // أ
var ALEF_AR_03 = []byte{216, 165} // إ
var ALEF_AR_04 = []byte{217, 177} // ٱ

var YE_FA = []byte{219, 140}    // ی
var YE_AR_01 = []byte{217, 138} // ي
var YE_AR_02 = []byte{216, 166} // ئ
var YE_AR_03 = []byte{217, 137} // ى

var VAV_FA = []byte{217, 136}    // و
var VAV_AR_01 = []byte{216, 164} // ؤ

var KAF_FA = []byte{218, 169}    // ک
var KAF_AR_01 = []byte{217, 131} // ك

// تابع اصلی برای تبدیل کاراکترهای عربی به فارسی
func arabicToPersianCharacter(text string) string {
	text = checkALEF(text)
	text = checkYE(text)
	text = checkVAV(text)
	text = checkKAF(text)
	return text
}

// تابع برای چک کردن ALEF و تبدیل آن به فارسی
func checkALEF(text string) string {
	text = strings.ReplaceAll(text, string(ALEF_AR_01), string(ALEF_FA))
	text = strings.ReplaceAll(text, string(ALEF_AR_02), string(ALEF_FA))
	text = strings.ReplaceAll(text, string(ALEF_AR_03), string(ALEF_FA))
	text = strings.ReplaceAll(text, string(ALEF_AR_04), string(ALEF_FA))
	return text
}

// تابع برای چک کردن YE و تبدیل آن به فارسی
func checkYE(text string) string {
	text = strings.ReplaceAll(text, string(YE_AR_01), string(YE_FA))
	text = strings.ReplaceAll(text, string(YE_AR_02), string(YE_FA))
	text = strings.ReplaceAll(text, string(YE_AR_03), string(YE_FA))
	return text
}

// تابع برای چک کردن VAV و تبدیل آن به فارسی
func checkVAV(text string) string {
	text = strings.ReplaceAll(text, string(VAV_AR_01), string(VAV_FA))
	return text
}

// تابع برای چک کردن KAF و تبدیل آن به فارسی
func checkKAF(text string) string {
	text = strings.ReplaceAll(text, string(KAF_AR_01), string(KAF_FA))
	return text
}

func main() {
	// تست با یک جمله عربی که می‌خواهیم به فارسی تبدیل کنیم
	text := "أريد كاف و ياء و واء في النص"
	convertedText := arabicToPersianCharacter(text)

	// نمایش نتیجه
	fmt.Println("Converted Text: ", convertedText)
}
