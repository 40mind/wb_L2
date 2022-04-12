package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var timeout = flag.Int("timeout", 10, "таймаут на подключение к серверу")
	log.SetFlags(0)
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		panic("Недостаточно аргументов")
	}

	host := args[0]
	port := args[1]

	start := time.Now()
	var conn net.Conn
	var err error

	for {
		if time.Since(start) >= (time.Duration(*timeout) * time.Second) {
			conn, err = net.Dial("tcp", host+":"+port)
			if err != nil {
				panic(err)
			}
			break
		}
	}
	defer conn.Close()
	log.Println("Connected")

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("Получено сообщение:\n" + message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if _, err := fmt.Fprintf(conn, scanner.Text()+"\n"); err != nil {
			log.Fatal("Закрыто")
		}
	}
}
