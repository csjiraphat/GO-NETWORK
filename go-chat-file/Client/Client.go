package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// Function sendFile
func sendFile(conn net.Conn, filePath string) {
	defer conn.Close() // close connection before exit

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // close file before exit

	// send file name to server
	_, fileName := filepath.Split(filePath)

	// Send file name
	conn.Write([]byte(fileName)) // send file name to server

	//create buffer for reading
	buffer := make([]byte, 1024)
	for {
		//read file to buffer
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File send complete")
			} else {
				fmt.Println("Error reading:", err)
			}
			return
		}
		//send buffer to server
		conn.Write(buffer[:n])
	}
}

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close() // close connection before exit
	//print message connect successfully
	fmt.Println("Connected to server successfully")

	reader := bufio.NewReader(os.Stdin)
	//get file path and file name from user
	fmt.Print("Enter file path+name: ")
	filepath, _ := reader.ReadString('\n')
	filepath = strings.TrimSpace(filepath)

	//send file to server
	sendFile(conn, filepath)
}
