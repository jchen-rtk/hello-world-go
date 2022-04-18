package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const msgDefault = "Hello"
const nameDefault = "World"

func main() {
	msg := msgDefault
	name := nameDefault

	// Parse environment variables.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := pair[0]
		value := pair[1]

		if key == "HELLO_MSG" {
			msg = value
		} else if key == "HELLO_NAME" {
			name = value
		}
	}

	// Parse command line arguments.
	if len(os.Args) == 2 {
		name = os.Args[1]
	}

	// Parse flags.
	nameFlag := flag.String("name", "", "Name to greet")
	flag.Parse()
	if *nameFlag != "" {
		name = *nameFlag
	}

	// Print to stdout.
	var output string
	msgSlice := strings.Split(msg, " ")
	if len(msgSlice) == 1 {
		output = msg + ", " + name + "!"
	} else {
		output = msg + " " + name
	}
	fmt.Println(output)
}
