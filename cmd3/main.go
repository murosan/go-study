package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	readout := exec.Command("./readout")
	stdin, _ := readout.StdinPipe()
	stdout, _ := readout.StdoutPipe()
	scanner := bufio.NewScanner(stdout)

	scannerOS := bufio.NewScanner(os.Stdin)

	go func() {
		for scannerOS.Scan() {
			line := scannerOS.Text()
			io.WriteString(stdin, line)
			if line == "fin" {
				return
			}
		}
	}()

	go func() {
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()

	if err := readout.Run(); err != nil {
		panic(err)
	}
}
