package gethoroscope

import "testing"

func TestGetDailyHoroscope(t *testing.T) {
	h := GetDailyHoroscope("gemini", "today")
	if h.Date != "" && h.Horoscope != "" && h.Sunsign != "" {
		t.Error("TEST FAILED")
	}
}
