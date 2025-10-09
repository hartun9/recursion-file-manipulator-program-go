package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	action := os.Args[1]
	inputPath := os.Args[2]
	outputPath := os.Args[3]

	switch action {
	case "reverse":
		// fileInput の内容を反転させ、その結果をfileOutputに書き出す。
		fmt.Println("ファイルの中身を逆にします")
		fmt.Println("inputPath: ", inputPath)
		fmt.Println("outputPath: ", outputPath)
		reverse(inputPath, outputPath)
	default:
		panic("コマンドが存在しません")
	}
}

func reverse(inputPath string, outputPath string) {
}
