# go-study


コマンドメモ

```sh
# 1 シンプルなコマンド実行(http://ascii.jp/elem/000/001/459/1459279/)
go run ./cmd1/main.go

#2 コマンドの出力を取得
go build -o ./count ./output/main.go
go run ./cmd2/main.go

#3 標準入力をコマンドに渡して、出力を受け取る
go build -o ./readout ./inout/main.go
go run ./cmd3/main.go

# 4 echo server using websocket
cd websocket/
go run ./main.go # access to http://localhost:8080/index.html
```
