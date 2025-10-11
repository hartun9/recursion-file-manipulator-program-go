package main

import (
	"fmt"
	"os"
	"strconv"
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
	case "copy":
		copy(os.Args[2], os.Args[3])
	case "duplicate-contents":
		n, err := strconv.Atoi(os.Args[4])
		if err != nil {
			panic(err)
		}
		deplicateContents(os.Args[2], n)
	case "replace-string":
		replaceString(os.Args[2], os.Args[3], os.Args[4])
	default:
		panic("コマンドが存在しません")
	}
}

func reverse(inputPath string, outputPath string) {
}

// copy inputpath にあるファイルのコピーを作成し、outputpath として保存する。
func copy(inputPath string, outputPath string) {
}

// deplicateContents inputpath にあるファイルの内容を読み込み、その内容を複製し、複製された内容を inputpath に n 回複製する。
func deplicateContents(inputPath string, n int) {
}

// replaceString inputpath にあるファイルの内容から文字列 'needle' を検索し、'needle' の全てを 'newstring' に置き換える。
func replaceString(inputPath string, needle string, newString string) {
}
