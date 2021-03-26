package services

import (
	"fmt"
	"io"
	"log"
	"os"
	"tcp-server/src/utils"
)

// SendFile sends file to client
func SendFile(filename, clientAddr string) []string {

	var result = make([]string, 1)
	const bufferSize = 1024

	file, erro := os.Open(filename)
	if erro != nil {
		log.Fatal(erro)
	}

	fileInfo, erro := file.Stat()
	if erro != nil {
		log.Fatal(erro)
	}
	fileName := utils.PrettyLoad(fileInfo.Name(), 64)

	log.Println(fmt.Sprintf("Sending %s file to client: %s", fileName, clientAddr))

	var buffer = make([]byte, bufferSize)

	log.Println(fmt.Sprintf("Starting file transfer to client %s", clientAddr))

	for {
		_, erro := file.Read(buffer)
		if erro == io.EOF {
			break
		}
	}

	responseText := fmt.Sprintf("File %s transfered successfully to client %s", fileName, clientAddr)
	log.Println(responseText)

	result = append(result, responseText)
	return result
}
