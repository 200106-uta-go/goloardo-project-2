package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/200106-uta-go/goloardo-project-2/pkg/models/servicecall"
)

// MyIP is the ip address of the host machine in which the program is running
var MyIP string

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		panic(err)
	}
	MyIP = addrs[(len(addrs) - 1)]
	fmt.Println("Host IP:", addrs[(len(addrs)-1)])
}

// SendNotify sends to the ark a notify command to register this service
func SendNotify(tipo string, arkip string, port string) {
	sc := servicecall.ServiceCall{
		Cmd:  "notify",
		IP:   MyIP,
		Tipo: tipo,
		Port: port,
	}

	body, err := json.Marshal(sc)
	if err != nil {
		fmt.Println("ERROR: not able to complete the HTTP POST to register this service")
		panic(err)
	}

	// Try not to hardcode ark's ip address
	resp, err := http.Post("http://"+arkip+":7777/notify", "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode == 503 {
		fmt.Println("ERROR: ark was not able to register this service ")
		panic(err)
	} else if resp.StatusCode == 201 {
		fmt.Println("Service register succesfully with response code", resp.Status)
	} else {
		panic("ERROR: Some other status code was returned beside 201(Created) and 503(Service Unavailable)")
	}
}

// SendVerify ...
func SendVerify(tipo string, arkip string) (string, string) {
	sc := servicecall.ServiceCall{
		Cmd:  "verify",
		IP:   "",
		Tipo: tipo,
		Port: "",
	}

	pbody, err := json.Marshal(sc)
	if err != nil {
		panic("ERROR: not able to complete the HTTP POST to verify the service this service")
	}

	// Try not to hardcode ark's ip address
	resp, err := http.Post("http://"+arkip+":7777/verify", "application/json", bytes.NewBuffer(pbody))
	if err != nil || resp.StatusCode == 503 {
		fmt.Println("ERROR: ark was not able to verify the service requested ")
		panic(err)
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(rbody, &sc)
	if err != nil {
		panic(err)
	}
	fmt.Println(sc)
	return sc.IP, sc.Port
}

// SendDestroy sends to the ark a destroy command to delete this service from the registered services
func SendDestroy(tipo string, arkip string, port string) {
	sc := servicecall.ServiceCall{
		Cmd:  "destroy",
		IP:   MyIP,
		Tipo: tipo,
		Port: port,
	}

	body, err := json.Marshal(sc)
	if err != nil {
		fmt.Println("NOTE: not able to complete the HTTP POST to destroy this service on the ark")
		return
	}

	// Try not to hardcode ark's ip address
	resp, err := http.Post("http://"+arkip+":7777/destroy", "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode == 503 {
		fmt.Println("NOTE: ark was not able to destroy this service", err)
	} else if resp.StatusCode == 202 {
		fmt.Println("Service destroy succesfully with response code", resp.Status)
	} else {
		fmt.Println("NOTE: Some other status code was returned beside 202(Accepted) and 503(Service Unavailable).")
		fmt.Println("This service was not destroy succesfully.")
	}
}
