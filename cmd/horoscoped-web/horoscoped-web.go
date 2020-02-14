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

func main() {

	itmpl := template.Must(template.ParseFiles("web/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := IndexContent{
			Horoscopes: gethoroscope.GetAllDailyHoroscope(),
		}
		itmpl.Execute(w, data)
	})

	println("Server is running on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Server crash:", err)
	}
}
