package myfile

import (
	"fmt"
	"os"
)

func CreatePath(path string) {
	var err = os.MkdirAll(path, 0750)

	if err != nil {
		fmt.Println(err)
	}
}

func PathExists(path string) bool {
	var _, err = os.Stat(path)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func ReadFile(path string) []byte {
	fileBuffer, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return fileBuffer
}

func WriteToFile(path string, content string) {
	os.WriteFile(path, []byte(content), 0644)
}
