package main

import (
	// "bufio"
	"fmt"
	"os"
	"os/exec"
	// "time"
)

func watch() error {
	fmt.Println("q:\tquit")
	fmt.Println("spc:\ttap")
	fmt.Println("\n")
	// https://stackoverflow.com/q/15159118
	//
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		str := string(b)

		// reader := bufio.NewReader(os.Stdin)
		// fmt.Printf("%s\n", "Enter text: ")
		// text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if str == "q" {
			break
		}
		fmt.Println("key:", str)
	}

	return nil
}
