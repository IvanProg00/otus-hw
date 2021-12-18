package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	if len(pflag.Args()) != 2 {
		log.Fatalln("Need 2 params")
	}
	hostname := pflag.Arg(0)
	port := pflag.Arg(1)
	NewTelnetClient(net.JoinHostPort(hostname, port), 10*time.Second, ioutil.NopCloser(os.Stdin), os.Stdout)
}
