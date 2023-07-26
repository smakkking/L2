package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

var (
	ErrTooManyArgument = errors.New("too many arguments")
)

func cmdCD(requst []string) {
	switch len(requst) {
	case 1:
		hm_dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Chdir(hm_dir)
	case 2:
		err := os.Chdir(requst[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		fmt.Fprintln(os.Stderr, ErrTooManyArgument)
	}
}

func cmdPWD(request []string) {
	if len(request) == 1 {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println(path)
		}
	} else {
		fmt.Fprintln(os.Stderr, ErrTooManyArgument)
	}
}

func cmdECHO(request []string) {
	for i := 1; i < len(request); i++ {
		fmt.Printf("%s ", request[i])
	}
	fmt.Println()
}

func cmdKILL(request []string) {
	switch len(request) {
	case 1:
		fmt.Fprintln(os.Stderr, "not enought argument")
	case 2:
		pid, err := strconv.Atoi(request[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		err = process.Kill()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	default:
		fmt.Fprintln(os.Stderr, ErrTooManyArgument)
	}
}

func cmdPS(request []string) {
	if len(request) != 1 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}
	sliceProc, _ := ps.Processes()

	for _, proc := range sliceProc {

		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())

	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		request := strings.Split(scanner.Text(), " ")
		switch request[0] {
		case "cd":
			cmdCD(request)
		case "pwd":
			cmdPWD(request)
		case "echo":
			cmdECHO(request)
		case "kill":
			cmdKILL(request)
		case "ps":
			cmdPS(request)
		}
	}
}
