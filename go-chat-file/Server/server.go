package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // close connection before exit
	//buffer for reading
	buffer := make([]byte, 1024)

	// Receive the file name from the client
	fileNameBuffer := make([]byte, 64) //64 bytes
	n, err := conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fileName := string(fileNameBuffer[:n])
	fmt.Println("Received file name:", fileName)

	//create file to store data
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // close file before exit

	// Recieve and write data to file
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File transfer complete")
				return
			} else {
				fmt.Println("Error reading:", err)
				return
			}
		}
		// Write data to file
		file.Write(buffer[:n])

	}
}

func main() {
	// Create Listener
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 5000") //print server is listening on port 5000

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		//print client connected
		fmt.Println("Client connected: ", conn.RemoteAddr())

		//Handle conection
		go handleConnection(conn) // Handle connection concurrently
	}
}
