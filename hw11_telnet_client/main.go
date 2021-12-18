package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var timeoutStr string
	pflag.StringVarP(&timeoutStr, "timeout", "t", "10s", "timeout")
	pflag.Parse()

	if len(pflag.Args()) != 2 {
		log.Fatal("Need 2 params")
	}

	hostname := pflag.Arg(0)
	port := pflag.Arg(1)
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("incorrect argument timeout: %v", err)
	}
	address := net.JoinHostPort(hostname, port)

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	ConnectTelnet(client, address, context.Background())
}

func ConnectTelnet(client TelnetClient, address string, context context.Context) {
	if err := client.Connect(); err != nil {
		log.Fatalf("connection failed: %s", err)
	}
	fmt.Fprintln(os.Stderr, "...Connected to", address)
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context, os.Interrupt)
	defer cancel()

	go func() {
		err := client.Receive()
		if err != nil {
			fmt.Println("receive error:", err)
		} else {
			fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		}
		cancel()
	}()
	go func() {
		err := client.Send()
		if err != nil {
			fmt.Println("send error:", err)
		} else {
			fmt.Fprintln(os.Stderr, "...EOF")
		}
		cancel()
	}()

	<-ctx.Done()
}
