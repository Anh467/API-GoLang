package control

import (
	"fmt"
	"io"
	"os"
)

func GetFile(path string) (*os.File, error) {
	configFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return nil, err
	}
	return configFile, nil
}

func ReadFile(file *os.File) {
	file.Seek(0, io.SeekStart)
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}

func GetFileContent(file *os.File) string {
	file.Seek(0, io.SeekStart)
	var content string = ""
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		content += string(buffer[:n])
		//fmt.Print(string(buffer[:n]))
	}
	return content
}
