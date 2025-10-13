package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// replaceString inputpath にあるファイルの内容から文字列 'needle' を検索し、'needle' の全てを 'newstring' に置き換える。
func replaceString(inputPath string, needle string, newString string) {
	fmt.Println("needleをnewStringに置き換えます")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("needle: ", needle)
	fmt.Println("newString: ", newString)

	file, err := os.OpenFile(inputPath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("content:")
	fmt.Println(string(content))

	newContent := strings.ReplaceAll(string(content), needle, newString)

	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

	if _, err := file.WriteString(newContent); err != nil {
		panic(err)
	}

	if err := file.Truncate(int64(len(newContent))); err != nil {
		panic(err)
	}
}
