package main

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/hamed-yousefi/gowl"
)

var (
	openvpnPath = "/usr/sbin/openvpn"
	serverDir   = "/etc/openvpn/server"
)

type proc struct {
	args  []string
	proto string
	port  int
	pid   string
}

func (p proc) Start() error {
	log.Println("starting "+p.Name(), "(", p.proto, ")", "port:", p.port)
	a := append(p.args, "--proto", p.proto, "--port", strconv.Itoa(p.port))
	log.Println("args:", a)
	cmd := exec.Command(openvpnPath, a...)
	cmd.Dir = serverDir
	/*
		out, err := cmd.Output()
		if err != nil {
			log.Println("error:", err)
			fmt.Println(string(out))
			return err
		}
		fmt.Println(out)
	*/
	cmd.Run()
	log.Println("cmd started")
	return nil
}

func (p proc) Name() string {
	return "openvpn-" + p.proto + "-" + strconv.Itoa(p.port)
}

func (p proc) PID() gowl.PID {
	return gowl.PID("p-" + p.pid)
}
