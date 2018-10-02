package main

import (
	"fmt"
	"time"
)

// 毎秒出力するだけ
// ビルドして cmd2/main.go から呼び、
// 標準出力をキャッチできるか確かめる
// go build -o count ./output/main.go

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
