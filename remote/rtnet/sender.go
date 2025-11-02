package rtnet

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type Sender struct {
	port int
	conn net.Conn

	healty bool
}

func NewSender(port int) (*Sender, error) {
	conn, err := dial(port)
	s := &Sender{
		port: port,
	}
	if err != nil {
		go s.dial()
	} else {
		s.conn = conn
		s.healty = true
	}
	return s, nil
}

func (s *Sender) dial() {
	for {
		conn, err := dial(s.port)
		if err != nil {
			log.Println("proba wdzwonienia zakonczona niepowodzeniem")
			time.Sleep(time.Second)
		} else {
			s.conn = conn
			break
		}
	}
	s.healty = true
}

func (s *Sender) Send(data []byte) error {
	if !s.healty {
		return errors.New("brak polaczenia z serwerem")
	}
	if _, err := s.conn.Write(append(data, '\n')); err != nil {
		log.Println("blad podczas wysylania danych")
	}

	msg, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil {
		if err == io.EOF {
			log.Println("stracono polaczenie z serwerem")
			s.conn.Close()
			for {
				s.conn, err = dial(s.port)
				if err != nil {
					log.Println("nieudana proba przywrocenia polaczenia")
					time.Sleep(time.Second)
				} else {
					break
				}
			}
			log.Println("ponownie polaczono z serwerem")
		}
	}
	log.Println("\nserver status: ", msg)
	return err
}

func dial(port int) (net.Conn, error) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Sender) Close() error {
	return s.conn.Close()
}
