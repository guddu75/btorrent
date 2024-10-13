package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"unicode"

	bencode "github.com/jackpal/bencode-go" // Available if you need it!
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
	} else if rune(bencodedString[0]) == 'd' {
		decodedDict, _, err := DecodeDict(bencodedString, 0)
		if err != nil {
			return "", err
		}
		return decodedDict, nil
	} else {
		return "", fmt.Errorf("Only strings are supported at the moment")
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	command := os.Args[1]

	if command == "decode" {
		bencodedValue := os.Args[2]

		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err, err.Error())
			return
		}
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else if command == "info" {
		data, err := os.ReadFile(os.Args[2])

		if err != nil {
			fmt.Println(err, err.Error())
		}

		decoded, _, err := DecodeDict(string(data), 0)

		if err != nil {
			fmt.Println(err, err.Error())
		}
		fmt.Println("Tracker URL:", decoded["announce"])

		info, ok := decoded["info"].(map[string]interface{})

		if info == nil || !ok {
			fmt.Println("no info section")
		}

		fmt.Println("Length:", info["length"])

		hasher := sha1.New()

		bencode.Marshal(hasher, info)

		fmt.Printf("Info Hash: %x\n", hasher.Sum(nil))

	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
