# Cypher Tool Guideline

# Overview

Cypher Tool is a simple command-line tool used to **encrypt || decrypt text** using different methods.

# Methods

1) **Reverse Cipher** – Reverses the input text.

  Ex. `hello` → `olleh`

2) **ROT13 Cipher** – Moves each letter 13 positions in the alphabet.

  Ex. `hello` → `uryyb`

3) **Morse Code** – Converts text to Morse code or Morse code back to text. Space are replaced with (/) when encrypting. 

  Ex. `SOS` → `... --- ...`
  
Note for the **Morse Code** , you must know morse code :))) 
If you dont we added the a guideline for each one of the letters at the end 
# How Does It Work?

1. Get input from the user.
2. Choose **Encrypt** or **Decrypt**.
3. Choose a method: **Reverse, ROT13, or Morse**.
4. Process the input.
5. Return the result to the user.

# How to Run

Make sure Go is installed, then run:

go run cyphertool.go


Follow the prompts shown in the terminal to choose the task, method, and input.

> **Note:** These methods are simple ciphers and are not suitable for securing sensitive information.


# MORSECODE TABLE 

	'a': ".-",   'b': "-...", 'c': "-.-.", 'd': "-..",
	'e': ".",    'f': "..-.", 'g': "--.",  'h': "....",
	'i': "..",   'j': ".---", 'k': "-.-",  'l': ".-..",
	'm': "--",   'n': "-.",   'o': "---",  'p': ".--.",
	'q': "--.-", 'r': ".-.",  's': "...",  't': "-",
	'u': "..-",  'v': "...-", 'w': ".--",  'x': "-..-",
	'y': "-.--", 'z': "--.." 
	
Example 
SOS → ... --- ...


# CopyRights ByteBreakers (Brent, Sami, Oday) 2026 
