package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"pl.home/remote/encoder"
	"pl.home/remote/message"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8981")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	text(conn)
	binary(conn)
}

func binary(conn net.Conn) {
	ball := message.BallMsg{
		X:     123,
		Y:     444,
		Speed: 8889,
	}

	encr := encoder.NewBinaryEncoder[message.BallMsg]()
	enc, err := encr.Encode(ball)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", enc)

	b, err := encr.Decode(enc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
	conn.Write(enc)
	conn.Write([]byte{'\n'})

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(msg)
	}
}

func text(conn net.Conn) {
	for {
		fmt.Print("wpisz wiadomosc: ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		_, err = conn.Write([]byte(text))
		fmt.Println("ERR: ", err)
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("stracono polaczenie z serwerem")
				conn.Close()
				conn, err = net.Dial("tcp", "127.0.0.1:8981")
				if err != nil {
					panic(err)
				}
				fmt.Println("ponownie polaczono z serwerem")
				conn.Write([]byte(text))
				status, err = bufio.NewReader(conn).ReadString('\n')
			}
		}
		fmt.Println("status: ", status, " err: ", err)
	}
}
