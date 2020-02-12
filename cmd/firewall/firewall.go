package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var flush string

func init() {
	flag.StringVar(&flush, "flush", "", "flush ip tables")
	flag.Parse()
}

func accept(port string) {
	err := iptables("-A", "INPUT", "-p", "tcp", "-s", "localhost", "--dport", port, "-j", "ACCEPT")

	if err != nil {
		log.Fatal(err)
	}
}

func drop(port string) {
	err := iptables("-A", "INPUT", "-p", "tcp", "--dport", port, "-j", "DROP")

	if err != nil {
		fmt.Println("YO")
		log.Fatal(err)
	}
}

func iptables(args ...string) error {
	cmd := exec.Command("iptables", args...)
	//fmt.Println("CMD", cmd)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("ERROR", string(out))
		log.Fatal(err)
	}

	return nil
}

func main() {

	if flush == "yes" {
		iptables("-F")
	} else {
		accept("8081")
		drop("8081")
	}

}
