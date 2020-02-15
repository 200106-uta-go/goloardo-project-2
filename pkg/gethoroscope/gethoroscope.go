package gethoroscope

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

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

//GetAllDailyHoroscope gets all signs horoscope information
func GetAllDailyHoroscope() (horoscopes []DailyHoroscope) {
	horoscopes = append(horoscopes,
		GetDailyHoroscope("Aries", "today"),
		GetDailyHoroscope("Taurus", "today"),
		GetDailyHoroscope("Gemini", "today"),
		GetDailyHoroscope("Cancer", "today"),
		GetDailyHoroscope("Leo", "today"),
		GetDailyHoroscope("Virgo", "today"),
		GetDailyHoroscope("Libra", "today"),
		GetDailyHoroscope("Scorpio", "today"),
		GetDailyHoroscope("Sagittarius", "today"),
		GetDailyHoroscope("Capricorn", "today"),
		GetDailyHoroscope("Aquarius", "today"),
		GetDailyHoroscope("Pisces", "today"),
	)

	return
}

//GetYearlyHoroscope configures JSON api to struct recieved and returns string including zodiac sign, year for horoscope, and horoscope based on user input
func GetYearlyHoroscope(userSunsign string, dateInput string) (userHoroscope YearlyHoroscope) {
	userSunsign = strings.Title(userSunsign)
	dateInput = strings.ToLower(dateInput)
	myURL := "http://horoscope-api.herokuapp.com/horoscope/" + dateInput + "/" + userSunsign
	response, _ := http.Get(myURL)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	json.Unmarshal(body, &userHoroscope)
	return
}

//GetAllYearlyHoroscope gets all signs horoscope information
func GetAllYearlyHoroscope() (horoscopes []YearlyHoroscope) {
	horoscopes = append(horoscopes,
		GetYearlyHoroscope("Aries", "year"),
		GetYearlyHoroscope("Taurus", "year"),
		GetYearlyHoroscope("Gemini", "year"),
		GetYearlyHoroscope("Cancer", "year"),
		GetYearlyHoroscope("Leo", "year"),
		GetYearlyHoroscope("Virgo", "year"),
		GetYearlyHoroscope("Libra", "year"),
		GetYearlyHoroscope("Scorpio", "year"),
		GetYearlyHoroscope("Sagittarius", "year"),
		GetYearlyHoroscope("Capricorn", "year"),
		GetYearlyHoroscope("Aquarius", "year"),
		GetYearlyHoroscope("Pisces", "year"),
	)

	return
}

//GetMonthlyHoroscope configures JSON api to struct recieved and returns YearlyHorsocope struct that includes zodiac sign, month for horoscope, and horoscope based on user input
func GetMonthlyHoroscope(userSunsign string, dateInput string) (userHoroscope MonthlyHoroscope) {
	userSunsign = strings.Title(userSunsign)
	dateInput = strings.ToLower(dateInput)
	myURL := "http://horoscope-api.herokuapp.com/horoscope/" + dateInput + "/" + userSunsign
	response, _ := http.Get(myURL)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	json.Unmarshal(body, &userHoroscope)
	return
}

//GetAllMonthlyHoroscope gets all signs horoscope information
func GetAllMonthlyHoroscope() (horoscopes []MonthlyHoroscope) {
	horoscopes = append(horoscopes,
		GetMonthlyHoroscope("Aries", "month"),
		GetMonthlyHoroscope("Taurus", "month"),
		GetMonthlyHoroscope("Gemini", "month"),
		GetMonthlyHoroscope("Cancer", "month"),
		GetMonthlyHoroscope("Leo", "month"),
		GetMonthlyHoroscope("Virgo", "month"),
		GetMonthlyHoroscope("Libra", "month"),
		GetMonthlyHoroscope("Scorpio", "month"),
		GetMonthlyHoroscope("Sagittarius", "month"),
		GetMonthlyHoroscope("Capricorn", "month"),
		GetMonthlyHoroscope("Aquarius", "month"),
		GetMonthlyHoroscope("Pisces", "month"),
	)

	return
}
