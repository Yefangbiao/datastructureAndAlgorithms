package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type port int

func (p *port) String() string {
	return string(*p)
}

func (p *port) Set(s string) error {
	var iPort int
	_, err := fmt.Sscanf(s, "%d", &iPort)
	return err
}

func PortFlag(name string, value port, usage string) *port {
	f := value
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

var p = PortFlag("port", 8000, "the port")

func main() {
	flag.Parse()
	address := fmt.Sprintf("localhost:%v", *p)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
