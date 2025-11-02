package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"pl.home/remote/encoder"
	"pl.home/remote/message"
	"pl.home/remote/rtnet"
)

func main() {

	// text()
	binary()
}

func binary() {

	sdr, err := rtnet.NewSender(8981)
	if err != nil {
		panic(err)
	}
	id := 0
	for {
		id++
		ball := message.BallMsg{
			Id:    id,
			X:     123,
			Y:     444,
			Speed: 8889,
			Time:  time.Now(),
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
		fmt.Printf("wysylam: %+v\n", b)
		if err := sdr.Send(enc); err != nil {
			fmt.Println("problem podczas wyslania wiadomosci")
		}
		time.Sleep(time.Second)
	}
}

func text() {
	conn, err := net.Dial("tcp", ":8981")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
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
				for {
					conn, err = net.Dial("tcp", "127.0.0.1:8981")
					if err != nil {
						time.Sleep(time.Second)
					} else {
						break
					}
				}
				fmt.Println("ponownie polaczono z serwerem")
				conn.Write([]byte(text))
				status, err = bufio.NewReader(conn).ReadString('\n')
			}
		}
		fmt.Println("status: ", status, " err: ", err)
	}
}
