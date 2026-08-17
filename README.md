# Cypher Tool Guideline

# Overview

Cypher Tool is a simple command-line tool used to **encrypt || decrypt text** using different methods.

# Methods

1) **Reverse Cipher** – Reverses the input text by replacing each letter with the letter that appears the same distance from the end of the alphabet. For example, a becomes z, b becomes y, and c becomes x. This means each letter is transformed into its opposite letter in the alphabet, while keeping the same order of the text.

  Ex. `hello` → `svool`

2. **ROT13 Cipher** – Replaces each letter with the letter that is 13 positions further along in the alphabet. Since the alphabet has 26 letters, applying ROT13 twice returns the original text. For example, `a` becomes `n`, `b` becomes `o`, and `c` becomes `p`.

Ex. `hello` → `uryyb`


3. **Morse Code** – Converts text into a series of dots (`.`) and dashes (`-`) that represent each letter, number, or symbol. When encrypting, spaces between words are replaced with a slash (`/`) to make it clear where one word ends and another begins. It can also convert Morse code back into normal text.

Ex. `SOS` → `... --- ...`

  
Note for the **Morse Code** , you must know morse code when decrypting :))) 
If you dont know how,w we added the a guideline for each one of the letters at the end 

# How Does It Work?

1. Get input from the user.
2. Choose **Encrypt** or **Decrypt**.
3. Choose a method: **Reverse, ROT13, or Morse**.
4. Process the input.
5. Return the result to the user.

# How to Run

Make sure Go is installed, then run in Terminal:

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
