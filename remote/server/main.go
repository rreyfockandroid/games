package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8981")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("serwer czeka na polaczenie")

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("polaczono z klientem")

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("otrzymano: ", msg)
		conn.Write([]byte("ok\n"))
	}
}
