package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Тестовый процесс")
		time.Sleep(5 * time.Second)
	}
}
