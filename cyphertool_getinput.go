package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (toEncrypt bool, encoding string, message string){
	toEncrypt := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to our amazing Cypher Tool!\n")
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")
	input, _ := ToEncrypt.ReadString('\n')
	choice := strings.TrimSpace(input)
		if choice == "1" {
			fmt.Println("You chose to encrypt!")
			func encrypt()
		} else if choice == "2" {
			fmt.Println("You chose to decrypt!")
			func decrypt()
		} else {
			fmt.Println("Invalid choice.")
		} 
	
	encoding := bufio.NewReader(os.Stdin)
	fmt.Println("Select cypher (1/2):")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("3. Morse.")
	input, _ := encoding.ReadString('n')
	cypherchoice := strings.TrimSpace(input)
		if cypherchoice == "1" {
			fmt.Println("You chose ROT13")
			func rot13()
		} else if cypherchoice == "2" {
			fmt.Println("You chose Reverse")
			func reverse()
		} else if cypherchoice == "3" {
			fmt.Println("You chose Morse")
			func morse()
		} else {
			fmt.Println("Invalid choice.")
		}
}
