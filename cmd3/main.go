package main

import (
	"bufio"
	"fmt"
	"gopkg.in/pipe.v2"
	"io"
	"os"
	"os/exec"
	"sync"
)

type CommandWrapper struct {
	cmd *exec.Cmd

	// コマンドの pipe
	stdin  io.WriteCloser
	stdout io.ReadCloser

	//　実行したコマンドが出力した値を読み取る Scanner
	sc *bufio.Scanner

	// コマンドが終了したかどうか
	// done が true なら終了したことを表す
	mux    sync.Mutex
	doneCh chan bool
}

// コマンドが出力した文字列を読み取って出力する
func cmdOutputPrinter(s *CommandWrapper) {
	for s.sc.Scan() {
		if line := s.sc.Text(); line != "fin" {
			fmt.Println("[output] " + line)
		}
	}

	if err := s.sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

// コマンドが終了したら done を true にする
func terminator(s *CommandWrapper) {
	s.mux.Lock()
	s.cmd.Wait()
	fmt.Println("Received exit code")
	s.doneCh <- true
	s.mux.Unlock()
}

func newCommand() *CommandWrapper {
	cmd := exec.Command("./readout")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	return &CommandWrapper{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,

		sc:     bufio.NewScanner(stdout),
		doneCh: make(chan bool),
	}
}

func main() {
	cw := newCommand()
	if err := cw.cmd.Start(); err != nil {
		panic(err)
	}

	go cmdOutputPrinter(cw)
	go terminator(cw)

	go func() {
		<-cw.doneCh
		fmt.Println("Exit.")
		os.Exit(0)
	}()

	p := pipe.Line(
		pipe.Read(os.Stdin),
		pipe.Write(cw.stdin),
	)

	if err := pipe.Run(p); err != nil {
		panic(err)
	}

	if s, err := pipe.Output(p); err != nil {
		fmt.Println("error: '" + string(s) + "'")
		panic(err)
	}
}
