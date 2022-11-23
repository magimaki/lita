package main

import (
	"lita/cli"
)

// go build -ldflags="-s -w" -o lita main.go
// go build -ldflags="-s -w" -o lita main.go && upx -9 lita
func main() {
	cli.Execute()
}
