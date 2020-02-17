package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//YearlyHoroscope is a configuration struct for api response of yearly horoscope based on user sunsign
type YearlyHoroscope struct {
	Horoscope string `json:"horoscope"`
	Sunsign   string `json:"sunsign"`
	Year      string `json:"year"`
}

//MonthlyHoroscope is a configuration struct for api response of monthly horoscope based on user sunsign
type MonthlyHoroscope struct {
	Horoscope string `json:"horoscope"`
	Month     string `json:"month"`
	Sunsign   string `json:"sunsign"`
}

func main() {
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
		var yh YearlyHoroscope
		turl := "http://horoscope-api.herokuapp.com/horoscope/year/" + s
		resp, _ := http.Get(turl)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		json.Unmarshal(body, &yh)
		key := yh.Sunsign + "year"
		yh.Horoscope = strings.TrimLeft(yh.Horoscope, "[")
		yh.Horoscope = strings.TrimRight(yh.Horoscope, "]")
		yh.Horoscope = strings.Trim(yh.Horoscope, "\"")
		//fmt.Println("Adding key:", key, "value:", yh.Horoscope)
		resp, err := http.PostForm("http://127.0.1.1:8081/write", url.Values{
			"key":   {key},
			"value": {yh.Horoscope},
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status)
	}

	for _, s := range signs {
		var mh MonthlyHoroscope
		turl := "http://horoscope-api.herokuapp.com/horoscope/month/" + s
		resp, _ := http.Get(turl)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		json.Unmarshal(body, &mh)
		key := mh.Sunsign + "month"
		mh.Horoscope = strings.TrimLeft(mh.Horoscope, "[")
		mh.Horoscope = strings.TrimRight(mh.Horoscope, "]")
		mh.Horoscope = strings.Trim(mh.Horoscope, "\"")
		//fmt.Println("Adding key:", key, "value:", yh.Horoscope)
		resp, err := http.PostForm("http://127.0.1.1:8081/write", url.Values{
			"key":   {key},
			"value": {mh.Horoscope},
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status)
	}
}
