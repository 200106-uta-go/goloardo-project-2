package gethoroscope

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//ConfHoroscope stores requested horoscope type as a string after configuration
var ConfHoroscope string

//NumToMonth stores month numeric references as keys and month names as values corresponding to their numeric equivalent
var NumToMonth = map[string]string{
	"01": "January",
	"02": "February",
	"03": "March",
	"04": "April",
	"05": "May",
	"06": "June",
	"07": "July",
	"08": "August",
	"09": "September",
	"10": "October",
	"11": "November",
	"12": "December",
}

//DailyHoroscope is a configuration struct for api response of daily horoscope based on user sunsign
type DailyHoroscope struct {
	Date      string `json:"date"`
	Horoscope string `json:"horoscope"`
	Sunsign   string `json:"sunsign"`
}

//WeeklyHoroscope is a configuration struct for api response of weekly horoscope based on user sunsign
type WeeklyHoroscope struct {
	Horoscope string `json:"horoscope"`
	Sunsign   string `json:"sunsign"`
	Week      string `json:"week"`
}

//MonthlyHoroscope is a configuration struct for api response of monthly horoscope based on user sunsign
type MonthlyHoroscope struct {
	Horoscope string `json:"horoscope"`
	Month     string `json:"month"`
	Sunsign   string `json:"sunsign"`
}

//YearlyHoroscope is a configuration struct for api response of yearly horoscope based on user sunsign
type YearlyHoroscope struct {
	Horoscope string `json:"horoscope"`
	Sunsign   string `json:"sunsign"`
	Year      string `json:"year"`
}

//GetDailyHoroscope configures JSON api to struct recieved and returns string including zodiac sign, day for horoscope, and horoscope based on user input
func GetDailyHoroscope(userSunsign string, dateInput string) (userHoroscope DailyHoroscope) {
	userSunsign = strings.Title(userSunsign)
	dateInput = strings.ToLower(dateInput)
	myURL := "http://horoscope-api.herokuapp.com/horoscope/" + dateInput + "/" + userSunsign
	response, _ := http.Get(myURL)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	json.Unmarshal(body, &userHoroscope)
	return
}

//GetWeeklyHoroscope configures JSON api to struct recieved and returns string of zodiac sign, week for horoscope, and horoscope based on user input
func GetWeeklyHoroscope(jsonStruct []byte) string {
	var UserHoroscope WeeklyHoroscope
	err := json.Unmarshal(jsonStruct, &UserHoroscope)

	if err != nil {
		log.Fatal(err)
	}

	if UserHoroscope.Horoscope == "[]" {
		ConfHoroscope = ""
		return ConfHoroscope
	}

	retDate := strings.Split(UserHoroscope.Week, "-")
	fullDate := "Here is your reading for the Week of " + NumToMonth[retDate[1]] + " " + retDate[0] + ", " + retDate[2] + "to " + NumToMonth[retDate[4]] + " " + retDate[3] + ", " + retDate[5]
	ConfHoroscope = "Hi, " + UserHoroscope.Sunsign + "!" + "\n" + fullDate + "\n" + UserHoroscope.Horoscope

	return ConfHoroscope
}

//GetMonthlyHoroscope configures JSON api to struct recieved and returns string of zodiac sign, month for horoscope, and horoscope based on user input
func GetMonthlyHoroscope(jsonStruct []byte) string {
	var UserHoroscope MonthlyHoroscope
	var monthAbbrv = map[string]string{
		"Jan":  "January",
		"Feb":  "February",
		"Mar":  "March",
		"Apr":  "April",
		"May":  "May",
		"Jun":  "June",
		"Jul":  "July",
		"Aug":  "August",
		"Sept": "September",
		"Oct":  "October",
		"Nov":  "November",
		"Dec":  "December",
	}

	err := json.Unmarshal(jsonStruct, &UserHoroscope)

	if err != nil {
		log.Fatal(err)
	}

	if UserHoroscope.Horoscope == "[]" {
		ConfHoroscope = ""
		return ConfHoroscope
	}

	retDate := strings.Split(UserHoroscope.Month, " ")
	fullDate := "Here is your reading for the Month of " + monthAbbrv[retDate[0]] + " " + retDate[1]
	horoLen := len(UserHoroscope.Horoscope) - 2
	UserHoroscope.Horoscope = UserHoroscope.Horoscope[2:horoLen]
	ConfHoroscope = "Hi, " + UserHoroscope.Sunsign + "!" + "\n" + fullDate + "\n" + UserHoroscope.Horoscope
	return ConfHoroscope

}

//GetYearlyHoroscope configures JSON api to struct recieved and returns a string of zodiac sign, year for horoscope, and horoscope based on user input
func GetYearlyHoroscope(jsonStruct []byte) string {
	var UserHoroscope YearlyHoroscope
	err := json.Unmarshal(jsonStruct, &UserHoroscope)

	if err != nil {
		log.Fatal(err)
	}

	if UserHoroscope.Horoscope == "[]" {
		ConfHoroscope = ""
		return ConfHoroscope
	}

	fullDate := "Here is your reading for the Year " + UserHoroscope.Year
	ConfHoroscope = "Hi, " + UserHoroscope.Sunsign + "!" + "\n" + fullDate + "\n" + UserHoroscope.Horoscope
	return ConfHoroscope
}
