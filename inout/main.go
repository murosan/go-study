package main

import (
	"bufio"
	"fmt"
	"os"
)

// `fin` というテキストを受け取るまで
// 読み取ったテキストをそのまま出力する
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "fin" {
			fmt.Println("exit")
			return
		}

		fmt.Println("from command: " + line)
	}
}
