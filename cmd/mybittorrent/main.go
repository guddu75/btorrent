package main

import (
	"encoding/json"
	"fmt"
	"os"
	"unicode"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

// Ensures gofmt doesn't remove the "os" encoding/json import (feel free to remove this!)
var _ = json.Marshal

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345

func decodeBencode(bencodedString string) (interface{}, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {
		decodedString, _, err := DecodeString(bencodedString, 0)
		if err != nil {
			return "", err
		}
		return decodedString, nil
	} else if rune(bencodedString[0]) == 'i' {
		decodedInt, _, err := DecodeInt(bencodedString, 0)
		if err != nil {
			return "", err
		}
		return decodedInt, nil
	} else if rune(bencodedString[0]) == 'l' {
		decodedList, _, err := DecodeList(bencodedString, 0)
		if err != nil {
			return "", err
		}
		return decodedList, nil
	} else {
		return "", fmt.Errorf("Only strings are supported at the moment")
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	command := os.Args[1]

	if command == "decode" {
		// Uncomment this block to pass the first stage

		bencodedValue := os.Args[2]

		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
