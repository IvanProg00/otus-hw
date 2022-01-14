package main

import (
	"fmt"
	"os"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/cmd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
