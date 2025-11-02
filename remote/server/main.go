package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"

	"pl.home/remote/encoder"
	"pl.home/remote/message"
	"pl.home/remote/rtnet"
)

const (
	port = 8981
)

func main() {

	// text()
	binary()
}

func binary() {
	mess := make(chan []byte)
	lst, err := rtnet.NewListener(port, mess)
	if err != nil {
		panic(err)
	}
	for {
		msg := <-mess
		encr := encoder.NewBinaryEncoder[message.BallMsg]()
		ball, err := encr.Decode([]byte(msg))
		if err != nil {
			continue
			// panic(err)
		}
		fmt.Printf("BALL: %+v\n", ball)
	}
	lst.Close()
}

func text() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
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
