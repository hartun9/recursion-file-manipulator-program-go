package main

import "fmt"

// deplicateContents inputpath にあるファイルの内容を読み込み、その内容を複製し、複製された内容を inputpath に n 回複製する。
func deplicateContents(inputPath string, n int) {
	fmt.Println("ファイルの内容をn回複製します")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("n: ", n)
}
