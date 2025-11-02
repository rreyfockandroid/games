package remote

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
)

type Listener struct {
	ln     net.Listener
	conn   net.Conn
	broker chan []byte
}

func NewListener(port int, broker chan []byte) (*Listener, error) {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	conn, err := ln.Accept()
	if err != nil {
		return nil, err
	}

	l := &Listener{
		ln:     ln,
		conn:   conn,
		broker: broker,
	}
	l.run()
	return l, err
}

func (l *Listener) run() {
	go func() {
		for {
			msg, err := bufio.NewReader(l.conn).ReadString('\n')
			if err == io.EOF {
				log.Println("klient rozlaczyl sie")
				l.conn.Close()
				l.conn, err = l.ln.Accept()
				if err != nil {
					panic(err)
				}
				log.Println("ponownie polaczono z klientem")
				continue
			}
			l.broker <- []byte(msg)
			l.sendStatus()
		}
	}()
}

func (l *Listener) sendStatus() {
	if _, err := l.conn.Write([]byte("ok\n")); err != nil {
		log.Printf("blad podczas wyslania statusu: %v", err)
	}
}

func (l *Listener) Close() error {
	return l.ln.Close()
}
