package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
  )

  type Client struct {
      nickname string
      connection net.Conn
  }
  
  func handle(conn net.Conn) {
    defer conn.Close()
    s := bufio.NewScanner(conn)
    for s.Scan() {
      fmt.Fprintln(conn, s.Text())
    }
  }
  
  func main() {
    ln, err := net.Listen("tcp", ":3333")

    if err != nil {
      log.Fatal(err)
    }
    
    connections := []Client{}
    
    for {
      conn, err := ln.Accept()
      if err != nil {
        log.Fatal(err)
      }
      go handle(conn)
    }
  }
  