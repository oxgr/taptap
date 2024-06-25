package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type (
	Times = []time.Time
	Tempo = float32
)

const MAX_LEN_TIMES = 10

func watch() error {
	fmt.Println("q     \t: quit")
	fmt.Println("r     \t: reset")
	fmt.Println("<any> \t: tap")
	fmt.Printf("\n\n")

	printTempo(0)

	// https://stackoverflow.com/q/15159118
	//
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	b := make([]byte, 1)
	var times Times
loop:
	for {
		_, err := os.Stdin.Read(b)
		key := string(b)

		if err != nil {
			return err
		}
		switch key {
		case "q":
			tempo := getTempo(times)
			printTempo(tempo)
			break loop
		case "r":
			times = make(Times, 0)
			printReset()

		default:
			times = logTime(times, time.Now())
			tempo := getTempo(times)
			printTempo(tempo)

			// fmt.Println("key:", key)
		}
	}

	return nil
}

func print(msg string) error {
	fmt.Printf("|\t%-18s|\r", msg)
	return nil
}

func printReset() error {
	print("reset!")
	return nil
}

func printTempo(tempo float32) error {
	bpm_msg := "BPM:"
	print(fmt.Sprintf("%s %.2f", bpm_msg, tempo))
	return nil
}

func getTempo(times Times) Tempo {
	if len(times) < 2 {
		return 0
	}

	var totalMs int
	var prev int

	for i, t := range times {
		curr := int(t.UnixMilli())
		if i == 0 {
			prev = curr
			continue
		}
		diff := curr - prev
		totalMs += diff
		prev = curr
	}

	avg := float32(totalMs) / float32(len(times)-1)
	return intervalToBpm(avg)
}

func intervalToBpm(interval float32) float32 {
	msPerMin := float32(60000)
	beats := msPerMin / interval
	return beats
}

func logTime(times Times, t time.Time) Times {
	if len(times) >= MAX_LEN_TIMES {
		times = times[1:]
	}
	times = append(times, t)

	// fmt.Println("time: ", t.UnixMilli())
	// fmt.Println("times: ", times)

	return times
}
