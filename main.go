package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)

	action := os.Args[1]

	switch action {
	case "reverse":
		reverse(os.Args[2], os.Args[3])
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
