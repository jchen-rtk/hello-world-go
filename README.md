Write Enhanced Hello World In Go

Requirements:

Create a program, hello, which will do the following:
- When run as: `./hello` will output:  Hello, World!
- When run as `./hello Rob` will output: Hello, Rob! (reads command line arguments)
- When run as `HELLO_NAME=Sally ./hello` will output: Hello, Sally! (reads from environment variable HELLO_NAME)
- When run as `HELLO_MSG=Bonjour ./hello`  will output: Bonjour, World! (reads from environment variable HELLO_MSG to change the greeting).
- When run as `HELLO_MSG='Greetings & Salutations' HELLO_NAME=Tiffany ./hello` will output: Greetings & Salutations Tiffany

Uses zap.SugaredLogger to output formatted debug logs:
- If no environment or command line args are used; OR
- The value of the command line argument if one is present; OR
- The value of the relevant environment variables (HELLO_MSG, HELLO_NAME ) if supplied.

Hint: zap logging is found here: https://github.com/uber-go/zap
