package services

import (
	"fmt"
	"io"
	"log"
	"os"
	"tcp-server/src/utils"
)

// SendFile sends file to client
func SendFile(filename, clientAddr string) ([]string, error) {

	var result = make([]string, 1)
	const bufferSize = 1024

	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return []string{}, err
	}
	fileName := utils.PrettyLoad(fileInfo.Name(), 64)

	log.Printf("Sending %s file to client: %s", fileName, clientAddr)

	var buffer = make([]byte, bufferSize)

	log.Printf("Starting file transfer to client %s", clientAddr)

	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
	}

	responseText := fmt.Sprintf("File %s transfered successfully to client %s", fileName, clientAddr)
	log.Println(responseText)

	result = append(result, responseText)
	return result, nil
}
