package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/200106-uta-go/goloardo-project-2/config"
	"github.com/200106-uta-go/goloardo-project-2/pkg/models/servicecall"
	"github.com/google/uuid"
)

type registry struct {
	id   uuid.UUID
	ip   string
	tipo string
	port string
}

var ark []registry

func main() {

	http.HandleFunc("/verify", verify)

	http.HandleFunc("/notify", notify)

	http.HandleFunc("/destroy", destroy)

	errorChan := make(chan error, 5)
	go func() {
		fmt.Println("Listening on ports 7777 (http)")
		errorChan <- http.ListenAndServe(":7777", nil)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Fatalln(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			os.Exit(0)
		}
	}
}

// Recieves a http post request with json body that has the servicecall struct
func notify(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			goto fall
		} else {
			var sc servicecall.ServiceCall
			err = json.Unmarshal(body, &sc)
			if err != nil {
				goto fall
			}
			if sc.Cmd == "notify" {
				reg := registry{
					id:   uuid.New(),
					ip:   sc.IP,
					tipo: sc.Tipo,
					port: sc.Port,
				}
				ark = append(ark, reg)
				w.WriteHeader(201)
				fmt.Println(ark)
				break
			} else {
				goto fall
			}
		}
	fall:
		fallthrough
	default:
		w.WriteHeader(503)
	}
}

func verify(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			goto fall
		} else {
			var sc servicecall.ServiceCall
			err = json.Unmarshal(body, &sc)
			if err != nil {
				goto fall
			}
			if sc.Cmd == "verify" {
				// Look on the ark archive for services of type sc.tipo
				for _, item := range ark {
					if item.tipo == sc.Tipo {
						scresponse := servicecall.ServiceCall{
							Cmd:  "arkresponse",
							IP:   item.ip,
							Tipo: item.tipo,
							Port: item.port,
						}
						resp, err := json.Marshal(scresponse)
						if err != nil {
							goto fall
						}
						w.Header().Set("Content-Type", "application/json")
						w.Write(resp)
						break
					}
				}
				break
			} else {
				goto fall
			}
		}
	fall:
		fallthrough
	default:
		w.WriteHeader(503)
	}
}

func destroy(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			goto fall
		} else {
			var sc servicecall.ServiceCall
			err = json.Unmarshal(body, &sc)
			if err != nil {
				goto fall
			}
			if sc.Cmd == "destroy" {
				// Look on the ark archive for services of type sc.tipo
				for i, item := range ark {
					if item.ip == sc.IP && item.port == sc.Port && item.tipo == sc.Tipo {
						ark = append(ark[:i], ark[i+1:]...)
						if err != nil {
							goto fall
						}
						fmt.Println(ark)
						w.WriteHeader(202)
						break
					}
				}
				break
			} else {
				goto fall
			}
		}
	fall:
		fallthrough
	default:
		w.WriteHeader(503)
	}
}
