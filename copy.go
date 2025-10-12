package main

import (
	"fmt"
	"os"
)

// copy inputpath にあるファイルのコピーを作成し、outputpath として保存する。
func copy(inputPath string, outputPath string) {
	fmt.Println("ファイルをコピーします")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("outputPath: ", outputPath)

	file, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("inputPathの中身:")
	fmt.Println(string(file))

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	outputFile.WriteString(string(file))
}
