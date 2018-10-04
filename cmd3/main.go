package main

import (
	"bufio"
	"fmt"
	"gopkg.in/pipe.v2"
	"os"
	"os/exec"
)

// https://qiita.com/udzura/items/bc65456ecdaacb69d47f

func inputSupervisor(sc *bufio.Scanner) {
	for sc.Scan() {
		if line := sc.Text(); line != "fin" {
			fmt.Println("input: '" + line + "'")
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func terminator(cmd *exec.Cmd, term chan bool) {
	cmd.Wait()
	fmt.Println("Received exit code")
	term <- true
}

func main() {
	cmd := exec.Command("./readout")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	go inputSupervisor(scanner)

	term := make(chan bool, 1)
	go terminator(cmd, term)

	go func() {
		<-term
		fmt.Println("Exited.")
		os.Exit(0)
	}()

	p := pipe.Line(
		pipe.Read(os.Stdin),
		pipe.Write(stdin),
	)

	if err := pipe.Run(p); err != nil {
		panic(err)
	}

	fmt.Println(5)

	if s, err := pipe.Output(p); err != nil {
		fmt.Println("error: '" + string(s) + "'")
		panic(err)
	}

	fmt.Println(6)
}
