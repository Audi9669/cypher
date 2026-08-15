package main

import (
	"bufio"   // lets us read input line-by-line from the terminal
	"fmt"     // lets us print text to the screen
	"os"      // gives us access to os.Stdin (the keyboard/terminal input)
	"strings" // gives us string functions like TrimSpace
)

// getInput prompts the user for the three pieces of info needed to run
// a cipher: whether to encrypt or decrypt, which cipher to use, and the
// message itself.
func getInput() (toEncrypt bool, encoding string, message string) {
	// reader wraps os.Stdin so we can read a full line at once,
	// instead of one character at a time
	reader := bufio.NewReader(os.Stdin)

	// print the welcome banner
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println() // blank line, just for spacing/readability

	// ---- ask what operation to do ----
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")

	// read what the user typed (includes the newline character at the end)
	opChoice, _ := reader.ReadString('\n')
	// the _ throws away the error value on purpose, to keep this simple

	// remove the newline and any extra spaces from what they typed
	opChoice = strings.TrimSpace(opChoice)

	// if they typed "1", toEncrypt becomes true; otherwise false
	toEncrypt = opChoice == "1"

	fmt.Println() // blank line before next section

	// ---- ask which cipher to use ----
	fmt.Println("Select cypher (1/2):")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")

	cypherChoice, _ := reader.ReadString('\n')
	cypherChoice = strings.TrimSpace(cypherChoice)

	// convert their number choice into a real cipher name string,
	// since "1" and "2" alone don't mean anything to other functions
	switch cypherChoice {
	case "1":
		encoding = "rot13"
	case "2":
		encoding = "reverse"
	}

	fmt.Println() // blank line before next section

	// ---- ask for the actual message to encrypt/decrypt ----
	fmt.Println("Enter the message:")

	messageInput, _ := reader.ReadString('\n')
	message = strings.TrimSpace(messageInput) // clean it up the same way as before

	// send all three collected values back to whoever called getInput()
	return toEncrypt, encoding, message
}

func main() {
	// call getInput and store what it returns into three variables
	toEncrypt, encoding, message := getInput()

	// print them out so we can confirm getInput worked correctly
	fmt.Println("toEncrypt:", toEncrypt)
	fmt.Println("encoding:", encoding)
	fmt.Println("message:", message)
}
```