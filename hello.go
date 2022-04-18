package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
)

const msgDefault = "Hello"
const nameDefault = "World"

func main() {
	// Initialize zap logging
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// Default values to print.
	msg := msgDefault
	name := nameDefault

	// Parse environment variables.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := pair[0]
		value := pair[1]

		if key == "HELLO_MSG" {
			msg = value
			sugar.Debugf("env HELLO_MSG=%s", msg)
		} else if key == "HELLO_NAME" {
			name = value
			sugar.Debugf("env HELLO_NAME=%s", name)
		}
	}

	// Parse command line arguments.
	if len(os.Args) == 2 {
		name = os.Args[1]
		sugar.Debugf("args name=%s", name)
	}

	// Parse flags.
	nameFlag := flag.String("name", "", "Name to greet")
	flag.Parse()
	if *nameFlag != "" {
		name = *nameFlag
		sugar.Debugf("flag name=%s", name)
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
