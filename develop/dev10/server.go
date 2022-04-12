package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	exit := make(chan interface{})
	in := make(chan string)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-exit:
				close(in)
				return
			default:
				if msg, err := reader.ReadString('\n'); err == nil {
					in <- msg
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case msg := <-in:
				fmt.Fprint(conn, strings.ToUpper(msg))
			case <-exit:
				return
			}
		}
	}()

	<-signals
	close(exit)
}
