package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	action := os.Args[1]
	switch action {
	case "reverse":
		fmt.Println("ファイルの中身を逆にします")
	default:
		fmt.Println("コマンドが存在しません")
	}
}
