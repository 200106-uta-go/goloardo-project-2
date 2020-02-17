package gethoroscope

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
func GetYearlyHoroscope(userSunsign string, dbip string, dbport string) (yh YearlyHoroscope) {
	yh.Sunsign = userSunsign
	userSunsign = strings.ToLower(userSunsign)
	yh.Year = fmt.Sprint(time.Now().Year())

	myURL := "http://" + dbip + ":" + dbport + "/read?key=" + userSunsign + "year"
	response, err := http.Get(myURL)
	if err != nil {
		yh.Horoscope = fmt.Sprint(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		yh.Horoscope = fmt.Sprint(err)
		return
	}

	// Filling up the body
	yh.Horoscope = string(body)

	return
}

//GetAllYearlyHoroscope gets all signs horoscope information
func GetAllYearlyHoroscope(dbip string, dbport string) (horoscopes []YearlyHoroscope) {
	horoscopes = append(horoscopes,
		GetYearlyHoroscope("Aries", dbip, dbport),
		GetYearlyHoroscope("Taurus", dbip, dbport),
		GetYearlyHoroscope("Gemini", dbip, dbport),
		GetYearlyHoroscope("Cancer", dbip, dbport),
		GetYearlyHoroscope("Leo", dbip, dbport),
		GetYearlyHoroscope("Virgo", dbip, dbport),
		GetYearlyHoroscope("Libra", dbip, dbport),
		GetYearlyHoroscope("Scorpio", dbip, dbport),
		GetYearlyHoroscope("Sagittarius", dbip, dbport),
		GetYearlyHoroscope("Capricorn", dbip, dbport),
		GetYearlyHoroscope("Aquarius", dbip, dbport),
		GetYearlyHoroscope("Pisces", dbip, dbport),
	)

	return
}

//GetMonthlyHoroscope configures JSON api to struct recieved and returns YearlyHorsocope struct that includes zodiac sign, month for horoscope, and horoscope based on user input
func GetMonthlyHoroscope(userSunsign string, dbip string, dbport string) (mh MonthlyHoroscope) {
	mh.Sunsign = userSunsign
	userSunsign = strings.ToLower(userSunsign)
	mh.Month = fmt.Sprint(time.Now().Month())

	myURL := "http://" + dbip + ":" + dbport + "/read?key=" + userSunsign + "month"
	response, err := http.Get(myURL)
	if err != nil {
		mh.Horoscope = fmt.Sprint(err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		mh.Horoscope = fmt.Sprint(err)
		return
	}

	// Filling up the body request was succesful
	mh.Horoscope = string(body)

	return
}

//GetAllMonthlyHoroscope gets all signs horoscope information
func GetAllMonthlyHoroscope(dbip string, dbport string) (horoscopes []MonthlyHoroscope) {
	horoscopes = append(horoscopes,
		GetMonthlyHoroscope("Aries", dbip, dbport),
		GetMonthlyHoroscope("Taurus", dbip, dbport),
		GetMonthlyHoroscope("Gemini", dbip, dbport),
		GetMonthlyHoroscope("Cancer", dbip, dbport),
		GetMonthlyHoroscope("Leo", dbip, dbport),
		GetMonthlyHoroscope("Virgo", dbip, dbport),
		GetMonthlyHoroscope("Libra", dbip, dbport),
		GetMonthlyHoroscope("Scorpio", dbip, dbport),
		GetMonthlyHoroscope("Sagittarius", dbip, dbport),
		GetMonthlyHoroscope("Capricorn", dbip, dbport),
		GetMonthlyHoroscope("Aquarius", dbip, dbport),
		GetMonthlyHoroscope("Pisces", dbip, dbport),
	)

	return
}
