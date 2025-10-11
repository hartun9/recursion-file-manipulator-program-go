package main

import "fmt"

// replaceString inputpath にあるファイルの内容から文字列 'needle' を検索し、'needle' の全てを 'newstring' に置き換える。
func replaceString(inputPath string, needle string, newString string) {
	fmt.Println("needleをnewStringに置き換えます")
	fmt.Println("inputPath: ", inputPath)
	fmt.Println("needle: ", needle)
	fmt.Println("newString: ", newString)
}
