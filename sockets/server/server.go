package main

import (
	"net"
	"fmt"
)

func main(){
	server_socket, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server listening on localhost:8080")

	for {
		client_socket, err := server_socket.Accept()
		if err != nil {
			fmt.Println("Error accepting client connection:", err)
			continue
		}
		
		fmt.Println("Client connected:", client_socket.RemoteAddr())

		dataBuffer:= make([]byte, 1024)

		data, err := client_socket.Read(dataBuffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			client_socket.Close()
			break
		}

		decoded_data := string(dataBuffer[:data])
		fmt.Printf("Received message from client (%s): %s\n", client_socket.RemoteAddr(), decoded_data)

		response := "Hello from server! "
		client_socket.Write([]byte(response))

		defer client_socket.Close()
		fmt.Printf("Connection with client (%s) closed.\n", client_socket.RemoteAddr())
	}	
}