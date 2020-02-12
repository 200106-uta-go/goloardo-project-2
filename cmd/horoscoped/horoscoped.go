package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/danish287/project-1/internal/dbClient"
	"github.com/danish287/project-1/internal/gethoroscope"
)

func main() {
	println("Server is running on port 8081")
	currTime := time.Now()
	path := "logs/horoscoped/" + currTime.Format("01-02-2006") + ".log"
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)

	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var usrEmail = r.FormValue("myEmail")
		var usrPassword = r.FormValue("myPassword")
		//reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)

		fmt.Fprint(w, "\n\n ", usrEmail, usrPassword)
		dbEmail := dbClient.FindEmail(usrEmail)
		if dbEmail == usrEmail {
			//will redirectback to index page for user to input fields again
			fmt.Println("Hi")
			fmt.Println(r.URL.Path)
			fmt.Println("test")
			http.ServeFile(w, r, "./web/reject.html")
		} else {
			//dbClient.AddUser(usrName, usrEmail, usrPassword, usrSunsign)
			fmt.Println("ADD USER TO ACCOUNT")
			http.ServeFile(w, r, "./web/welcome.html")

		}

	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		var usrSunsign = r.FormValue("mySunsign")
		var usrName = r.FormValue("myName")
		var usrEmail = r.FormValue("myEmail")
		var usrPassword = r.FormValue("myPassword")
		answer := usrEmail + " " + usrName + " " + usrPassword + " " + usrSunsign
		fmt.Println(usrName)
		//reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)

		dbEmail := dbClient.FindEmail(usrEmail)
		if dbEmail == usrEmail {
			//will redirectback to index page for user to input fields again
			fmt.Println("Hi")
			fmt.Println(r.URL.Path)
			fmt.Println("test")
			http.ServeFile(w, r, "./web/reject.html")
		} else {
			//dbClient.AddUser(usrName, usrEmail, usrPassword, usrSunsign)
			fmt.Println("ADD USER TO ACCOUNT")
			http.ServeFile(w, r, "./web/welcome.html")

		}
		fmt.Println(answer)

	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println("test")
		http.ServeFile(w, r, "./web/welcome.html")

	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/index.html")
	})

	http.HandleFunc("/signupmvp", func(w http.ResponseWriter, r *http.Request) {
		var usrSunsign = r.FormValue("mySunsign")
		var usrName = r.FormValue("myName")
		var usrEmail = r.FormValue("myEmail")
		var usrPassword = r.FormValue("myPassword")
		dbEmail := dbClient.FindEmail(usrEmail)

		if dbEmail == usrEmail {
			http.ServeFile(w, r, "./web/reject.html")
		} else {
			dbClient.AddUser(usrName, usrEmail, usrPassword, usrSunsign, 0, false)
			reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, "today")
			fmt.Fprintf(w, reqHoroscope)
		}
	})

	http.HandleFunc("/loginmvp", func(w http.ResponseWriter, r *http.Request) {
		var usrEmail = r.FormValue("myEmail")
		var usrPassword = r.FormValue("myPassword")
		dbEmail := dbClient.FindEmail(usrEmail)

		if dbEmail != "" {
			isUser := dbClient.Auth(usrEmail, usrPassword)
			usrStatus := dbClient.IsBlocked(usrEmail)
			myUsr := dbClient.FindUsr(dbEmail)

			if usrStatus {

				yr, month, day := time.Now().Date()
				myTime := string(yr) + string(month) + string(day)
				fmt.Fprintf(w, "Too many failed password attempts. Due to security reasons, your account has been locked until tomorrow. Goodbye.")
				log.Output(0, (myTime + ": " + usrEmail + " - too many failed password attempts"))
			} else {
				if isUser {
					reqHoroscope := gethoroscope.GetHoroscope(myUsr[2], "today")
					fmt.Fprintf(w, reqHoroscope)
				} else {
					dbClient.FailedAttempt(dbEmail)
					if dbClient.IsBlocked(usrEmail) {
						yr, month, day := time.Now().Date()
						myTime := string(yr) + string(month) + string(day)
						fmt.Fprintf(w, "Too many failed password attempts. Due to security reasons, your account has been locked until tomorrow. Goodbye.")
						log.Output(0, (myTime + ": " + usrEmail + " - too many failed password attempts"))
					} else {
						http.ServeFile(w, r, "./web/index.html")
					}
				}
			}

		} else {
			http.ServeFile(w, r, "./web/account.html")
		}

	})

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}

}
