package filemanip

import (
	"fmt"
	"io"
	"os"
)

// DeplicateContents inputpath にあるファイルの内容を読み込み、その内容を複製し、複製された内容を inputpath に n 回複製する。
func DeplicateContents(inputPath string, n int) {
	fmt.Println("ファイルの内容をn回複製します")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("n: ", n)

	file, err := os.OpenFile(inputPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	fmt.Println("inputPathの中身:")
	fmt.Println(string(content))

	count := 0
	for count < n {
		_, err := file.WriteString(string(content))
		if err != nil {
			panic(err)
		}
		count++
	}
}
