package main

import (
	"bufio"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type TelnetClientConfig struct {
	address string
	timeout time.Duration
	net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {

	return &TelnetClientConfig{
		address: address,
		timeout: timeout,
	}
}

func (t *TelnetClientConfig) Connect() error {

	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}
	bufio.NewScanner(conn)
	t.Conn = conn
	return nil
}

func (t *TelnetClientConfig) Send() error {
	t.Conn.Write()

	return nil
}

func (t *TelnetClientConfig) Receive() error {
	return nil
}

func (t *TelnetClientConfig) Close() error {
	return t.Conn.Close()
}
