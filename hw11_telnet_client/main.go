package main

import (
	"io/ioutil"
	"os"
	"time"
)

func main() {
	hostname := os.Args[1]
	port := os.Args[2]

	NewTelnetClient(hostname+":"+port, 10*time.Second, ioutil.NopCloser(os.Stdin), os.Stdout)

	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
}
