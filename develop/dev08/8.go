package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/load"
	"os"
	exec2 "os/exec"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите команды")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Fields(line)[0]
		args := strings.Fields(line)

		switch command {
		case "cd":
			cd(args)
		case "pwd":
			path, _ := os.Getwd()
			fmt.Println(path)
		case "echo":
			echo(args)
		case "kill":
			kill(args)
		case "ps":
			miscStat, _ := load.Misc()
			fmt.Println(miscStat)
		case "fork":
			fork(args)
		case "exec":
			exec(args)
		case "quit":
			fmt.Println("Завершение программы")
			os.Exit(0)
		default:
			fmt.Println("Такой команды нет")
		}
	}
}

func cd(args []string) {
	if len(args) < 2 {
		fmt.Println("cd: Аргументов не хватает")
		return
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Println(err)
	}
}

func echo(args []string) {
	if len(args) < 2 {
		fmt.Println("echo: Аргументов не хватает")
		return
	}
	for i := 1; i < len(args); i++ {
		fmt.Printf("%s ", args[i])
	}
	fmt.Println()
}

func kill(args []string) {
	if len(args) < 2 {
		fmt.Println("kill: Аргументов не хватает")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("kill: неверный аргумент process id (нужно int)")
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("kill: %s\n", err)
		return
	}
	proc.Kill()
}

func fork(args []string) {
	if len(args) < 2 {
		fmt.Println("fork: Аргументов не хватает")
		return
	}

	cmd := exec2.Command(args[1])
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func exec(args []string) {
	if len(args) < 2 {
		fmt.Println("exec: Аргументов не хватает")
		return
	}

	content, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("exec: Не удалось открыть файл")
		return
	}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Fields(line)[0]
		args := strings.Fields(line)

		switch command {
		case "cd":
			cd(args)
		case "pwd":
			path, _ := os.Getwd()
			fmt.Println(path)
		case "echo":
			echo(args)
		case "kill":
			kill(args)
		case "ps":
			miscStat, _ := load.Misc()
			fmt.Println(miscStat)
		case "quit":
			fmt.Println("Завершение программы")
			os.Exit(0)
		default:
			fmt.Println("Такой команды нет")
		}
	}
}
