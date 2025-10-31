package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8981")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		fmt.Print("wpisz wiadomosc: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(text))

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(msg)

	}
}
