package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"pl.home/remote/encoder"
	"pl.home/remote/message"
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
	text(ln, conn)
	// binary(ln, conn)
}

func binary(ln net.Listener, conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// break
			}
		}
		encr := encoder.NewBinaryEncoder[message.BallMsg]()
		ball, err := encr.Decode([]byte(msg))
		if err != nil {
			// panic(err)
		}
		fmt.Printf("%+v\n", ball)
		conn.Write([]byte("ok\n"))
	}
}

func text(ln net.Listener, conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("ERR: ", err)
		if err == io.EOF {
			fmt.Println("klient rozlaczyl sie")

			conn.Close()
			conn, err = ln.Accept()
			fmt.Println("ponownie polaczono z klientem")
			if err != nil {
				panic(err)
			}
			// continue
		}
		fmt.Print("otrzymano: ", msg)
		_, err = conn.Write([]byte("ok\n"))
		fmt.Println("status error: ", err)
	}
}
