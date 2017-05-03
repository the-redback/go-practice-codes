package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main() {

	filepath.Walk(".",func(path string,info os.FileInfo,err error) error{
		fmt.Println(path)
		return nil
	})
}
