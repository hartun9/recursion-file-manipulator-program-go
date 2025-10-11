package main

import "fmt"

// copy inputpath にあるファイルのコピーを作成し、outputpath として保存する。
func copy(inputPath string, outputPath string) {
	fmt.Println("ファイルをコピーします")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("outputPath: ", outputPath)
}
