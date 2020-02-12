package config

import (
	"flag"
	"strings"
	"time"
)

//UserSunsign stores zodiac sign of user based on date of birth using the 12 traditionally defined zodiac signs
var UserSunsign string

//RequestedDate stores time of the requested reading. Values: today, week, month, and year
var RequestedDate string

//UserName contains the name of the user to store horoscope
var UserName string

//StoreHoroscope is a string refering to weather we will store the user's given horoscope or not
var StoreHoroscope string

//FileName stores the the name of the file based on user input to save if user would like to save a horoscope in a text file
var FileName string

func init() {
	flag.StringVar(&UserSunsign, "s", "", "User's sunsign")
	flag.StringVar(&RequestedDate, "date", "today", "Time for requested reading.\n   Values: Today, Week, Month, and Year")
	flag.Parse()

	currentTime := strings.Split(time.Now().String(), " ")[0]
	FileName = strings.ToLower(UserSunsign) + strings.Title(RequestedDate) + currentTime
}
