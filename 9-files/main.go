package main

import (
	"fmt"
	"os"

	fileManagement "test.com/go/files/file_management"
)

func main() {
	rootPath, _ := os.Getwd()
	data, err := fileManagement.ReadTextFile(rootPath + "/data/text.txt")
	if err != nil {
		panic(err)
	}
	fileManagement.WriteToFile(rootPath+"/data/output.txt", data)
	fmt.Println(data)
}
