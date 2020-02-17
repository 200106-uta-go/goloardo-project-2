package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//YearlyHoroscope is a configuration struct for api response of yearly horoscope based on user sunsign
type YearlyHoroscope struct {
	Horoscope string `json:"horoscope"`
	Sunsign   string `json:"sunsign"`
	Year      string `json:"year"`
}

func main() {

	var yh YearlyHoroscope
	signs := [12]string{
		"aries",
		"taurus",
		"gemini",
		"cancer",
		"leo",
		"virgo",
		"libra",
		"scorpio",
		"sagittarius",
		"capricorn",
		"aquarius",
		"pisces",
	}

	for _, s := range signs {
		url := "http://horoscope-api.herokuapp.com/horoscope/year/" + s
		resp, _ := http.Get(url)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		json.Unmarshal(body, &yh)
		key := yh.Sunsign + "year"
		//fmt.Println("Adding key:", key, "value:", yh.Horoscope)
		resp, err := http.Get("http://127.0.1.1:8081/write?key=" + key + "&value=" + yh.Horoscope)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status)
	}
}
