package gethoroscope

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidArgs(t *testing.T) {
	validArgs := GetHoroscope("aquarius", "today")
	isEmpty := strings.Contains(validArgs, "Aquarius")
	if !isEmpty {
		t.Error("TEST FAILED: valid parameters")
	}
}

func TestValidDailyResponse(t *testing.T) {
	validArgs := GetHoroscope("aquarius", "today")
	validDate := strings.Contains(validArgs, "Day")
	validSign := strings.Contains(validArgs, "Aquarius")
	if !validDate && !validSign {
		t.Error("TEST FAILED: valid daily response")
	}
}

func TestValidWeeklyResponse(t *testing.T) {
	validArgs := GetHoroscope("aquarius", "week")
	validDate := strings.Contains(validArgs, "Week")
	validSign := strings.Contains(validArgs, "Aquarius")
	if !validDate && !validSign {
		t.Error("TEST FAILED: valid weekly response")
	}
}

func TestValidMonthlyResponse(t *testing.T) {
	validArgs := GetHoroscope("aquarius", "month")
	validDate := strings.Contains(validArgs, "Month")
	validSign := strings.Contains(validArgs, "Aquarius")
	if !validDate && !validSign {
		t.Error("TEST FAILED: valid weekly response")
	}
}

func TestValidYearlyResponse(t *testing.T) {
	validArgs := GetHoroscope("aquarius", "year")
	validDate := strings.Contains(validArgs, "Year")
	validSign := strings.Contains(validArgs, "Aquarius")
	if !validDate && !validSign {
		t.Error("TEST FAILED: valid weekly response")
	}
}

func TestInvalidSunsign(t *testing.T) {
	invalidArgs := GetHoroscope("invalidSunSign", "today")
	isError := strings.Contains(invalidArgs, "Please try again using a valid sunsign")
	if !isError {
		t.Error("TEST FAILED: invalid sunsign")
	}
}

func TestInvalidDate(t *testing.T) {
	invalidArgs := GetHoroscope("libra", "nowhere")
	isError := strings.Contains(invalidArgs, "Please try again using a valid date")

	if !isError {
		t.Error("TEST FAILED: invalid date")
	}
}

func ExampleGetHoroscope() {
	myHoroscope := GetHoroscope("aries", "today")
	fmt.Println(myHoroscope)
}
