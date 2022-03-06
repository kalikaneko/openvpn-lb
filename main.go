package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hamed-yousefi/gowl"
)

var (
	nCPU        = 1
	tcpMode     = true
	defaultPort = 1194
	monitorAddr = "127.0.0.1:8080"
)

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "just a random server")
}

func main() {
	log.Println("Starting openvpn-lb")
	args := os.Args[1:]

	envCPU := os.Getenv("NCPU")
	if envCPU != "" {
		n, err := strconv.Atoi(envCPU)
		if err == nil {
			nCPU = n
		}
	}

	log.Println("Args:", args)
	pool := gowl.NewPool(nCPU * 2)
	for i := defaultPort; i < defaultPort+nCPU; i++ {
		u := proc{args, "udp4", i, strconv.Itoa(i) + "u"}
		pool.Register(u)
		if tcpMode {
			t := proc{args, "tcp4", i, strconv.Itoa(i) + "t"}
			pool.Register(t)
		}
	}
	pool.Start()
	log.Println("listening on", monitorAddr)
	log.Fatal(http.ListenAndServe(monitorAddr, nil))
}
