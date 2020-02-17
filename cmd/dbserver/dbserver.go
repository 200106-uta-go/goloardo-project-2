package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/200106-uta-go/goloardo-project-2/config"
	"github.com/200106-uta-go/goloardo-project-2/pkg/dbutil"
	"github.com/dgraph-io/badger"
)

// Opts defines settings for the database structure
type Opts struct {
	Directory string `json:"DataDirectory"`
}

var arkip string
var myport string

func main() {
	flag.StringVar(&arkip, "ark", "127.0.0.1", "This flag is used to specify the ip of the arkcontroller. DEFAULT = 127.0.0.1")
	flag.StringVar(&myport, "p", "8081", "This flag is used to specify the app port. DEFAULT = 8081")

	// Open a badger database with the defined directory from 'opts'
	db, err := badger.Open(badger.DefaultOptions("./app/badger"))
	if err != nil {
		panic(err)
	}
	// Defer the closing of our database so that we can acess it later on.
	defer db.Close()

	config.SendNotify("db", myport, arkip)
	fmt.Println("Hosting badger database server on port:", myport)

	// Create a multiplexer to host mutliple endpoints in one struct
	mux1 := http.NewServeMux()
	// Index is defaulted to respond with the port.
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "port: %s", myport)
	})
	// Calls a read command to the database.
	mux1.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		//fmt.Println(key)
		value := dbutil.DbRead(db, key)
		fmt.Println(value)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%v", value)
	})
	// Calls a write command to the database.
	mux1.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			//fmt.Fprintf(w, "ParseForm() err: %v", err)
			fmt.Println(w, "ParseForm() err: %v", err)
			return
		}
		key := r.FormValue("key")
		value := r.FormValue("value")
		fmt.Println(key, value)
		rKey, rVal := dbutil.DbWrite(db, key, value)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, rKey, rVal)
	})

	// ListenAnd Serve the multiplexer functions on designated port and send the error to a capture channel
	// , the capture channel is waiting for an error to handle it below in the for loop
	errorChan := make(chan error, 5)
	go func() {
		errorChan <- http.ListenAndServe(":"+myport, mux1)
	}()

	// Used to capture the sigint(ctr+c) to print before exiting
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				config.SendDestroy("db", arkip, myport)
				panic(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			config.SendDestroy("db", arkip, myport)
			os.Exit(0)
		}
	}
}
