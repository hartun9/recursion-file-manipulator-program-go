package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	action := os.Args[1]
	fileInput := os.Args[2]
	fileOutput := os.Args[3]

	switch action {
	case "reverse":
		fmt.Println("ファイルの中身を逆にします")
		fmt.Println("fileInput: ", fileInput)
		fmt.Println("fileOutput: ", fileOutput)
	default:
		panic("コマンドが存在しません")
	}
}
