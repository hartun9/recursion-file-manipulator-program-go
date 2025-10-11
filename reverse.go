package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// reverse fileInput の内容を反転させ、その結果をfileOutputに書き出す。
func reverse(inputPath string, outputPath string) {
	fmt.Println("ファイルの中身を逆にします")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("outputPath: ", outputPath)
	file, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("\ninputPathの中身:")
	fmt.Println(string(file))

	lines := strings.Split(string(file), "\n")
	fmt.Println(lines)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	slices.Reverse(lines)
	for i, line := range lines {
		if i == 0 && line == "" {
			continue
		}
		outputFile.WriteString(line + "\n")
	}
}
