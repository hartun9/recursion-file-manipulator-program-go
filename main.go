package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hartun9/recursion-file-manipulator-program-go/internal/filemanip"
)

func main() {
	fmt.Println(os.Args)

	action := os.Args[1]

	switch action {
	case "reverse":
		filemanip.Reverse(os.Args[2], os.Args[3])
	case "copy":
		filemanip.Copy(os.Args[2], os.Args[3])
	case "duplicate-contents":
		n, err := strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}
		filemanip.DeplicateContents(os.Args[2], n)
	case "replace-string":
		filemanip.ReplaceString(os.Args[2], os.Args[3], os.Args[4])
	default:
		panic("コマンドが存在しません")
	}
}
