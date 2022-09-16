package main

import (
	"embed"
)

var fileString string

var fileByte []byte

var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("foler/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file.hash")
	print(string(content2))
}