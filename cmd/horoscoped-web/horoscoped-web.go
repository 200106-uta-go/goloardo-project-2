package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/200106-uta-go/goloardo-project-2/config"
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

var dbip string
var dbport string

func main() {
	config.SendNotify("horowebserver", "8080")
	dbip, dbport = config.SendVerify("db")

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
			YearlyHoros: gethoroscope.GetAllYearlyHoroscope(dbip, dbport),
		}
		tmpl2.Execute(w, data)
	})

	http.HandleFunc("/month", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := MonthlyContent{
			MonthlyHoros: gethoroscope.GetAllMonthlyHoroscope(dbip, dbport),
		}
		tmpl3.Execute(w, data)
	})

	// Start server & Setup channels
	fmt.Println("Horoscope server is serving at port 8080...")
	errorChan := make(chan error, 2)
	go func() {
		errorChan <- http.ListenAndServe(":8080", nil)
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	for {
		select {
		case err := <-errorChan:
			if err != nil {
				// Bottom method sends the destroy signal to the ark
				config.SendDestroy("horowebserver", "8080")
				log.Fatalln(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			// Bottom method sends the destroy signal to the ark
			config.SendDestroy("horowebserver", "8080")
			os.Exit(0)
		}
	}
}
