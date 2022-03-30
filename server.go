// Goroutines are functions or methods that run concurrently with other 
// functions or methods. Goroutines can be thought of as lightweight threads. 
// The cost of creating a Goroutine is tiny when compared to a thread. Hence 
// it's common for Go applications to have thousands of Goroutines running concurrently. 

package main

import (
	"log"
	"net"
)

func read_client_command(client_command chan string) {
	command <- client_command

}

func read_server_command(srvr_command chan string) {}

func main() {
  conn, err := net.Dial("tcp", "localhost:1234")
  if err != nil {
    log.Fatal(err)
  }

  defer conn.Close()
	

}