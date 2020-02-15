package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/200106-uta-go/goloardo-project-2/pkg/gethoroscope"
)

// IndexContent ...
type IndexContent struct {
	Horoscopes []gethoroscope.DailyHoroscope
}

//YearlyContent stores all yearly horoscopes for each sunsign to populate page
type YearlyContent struct {
	YearlyHoros []gethoroscope.YearlyHoroscope
}

//MonthlyContent stores all monthly horoscopes for each sunsign to populate page
type MonthlyContent struct {
	MonthlyHoros []gethoroscope.MonthlyHoroscope
}

func main() {

	itmpl := template.Must(template.ParseFiles("web/index.html"))
	tmpl2 := template.Must(template.ParseFiles("web/yearly.html"))
	tmpl3 := template.Must(template.ParseFiles("web/monthly.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := IndexContent{
			Horoscopes: gethoroscope.GetAllDailyHoroscope(),
		}
		itmpl.Execute(w, data)
	})

	http.HandleFunc("/year", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := YearlyContent{
			YearlyHoros: gethoroscope.GetAllYearlyHoroscope(),
		}
		tmpl2.Execute(w, data)
	})

	http.HandleFunc("/month", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := MonthlyContent{
			MonthlyHoros: gethoroscope.GetAllMonthlyHoroscope(),
		}
		tmpl3.Execute(w, data)
	})

	println("Server is running on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Server crash:", err)
	}
}
