package main

import (
	"net"
	"fmt"
)
func main() {
	
	
	for {
		client_socket, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return
		}

		var message string
		fmt.Print(">>>")
		fmt.Scanln(&message)

		client_socket.Write([]byte(message))
		
		responseBuffer := make([]byte, 1024)
		response_decoded, err := client_socket.Read(responseBuffer)

		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		fmt.Printf("Response from server: %s\n", string(responseBuffer[:response_decoded]))
		client_socket.Close()
	}
}