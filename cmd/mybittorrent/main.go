package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
	// Available if you need it!
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
	} else if command == "info" || command == "peers" {
		data, err := os.ReadFile(os.Args[2])

		if err != nil {
			fmt.Println(err, err.Error())
		}

		decoded, _, err := DecodeDict(string(data), 0)

		if err != nil {
			fmt.Println(err, err.Error())
		}
		// fmt.Println("Tracker URL:", decoded["announce"])

		info, ok := decoded["info"].(map[string]interface{})

		if info == nil || !ok {
			fmt.Println("no info section")
		}

		if command == "info" {
			hash, err := getHash(info)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Print("Tracker URL: ", decoded["announce"])
			fmt.Print("Length: ", info["length"])
			fmt.Printf("Info Hash: %x", hash)
			fmt.Print("Piece Length: ", info["piece length"])
			fmt.Printf("Piece Hashes: %x", info["pieces"])
		} else if command == "peers" {
			resp, err := getPeers(decoded)
			if err != nil {
				fmt.Println("Response not happened error", err.Error())
			}
			decodedResp, _, err := DecodeDict(resp, 0)
			if err != nil {
				fmt.Println(err.Error())
			}
			// fmt.Println("decodedResp", decodedResp)
			peers := decodedResp["peers"].(string)
			for i := 0; i < len(peers); i += 6 {
				port := strconv.Itoa(int(binary.BigEndian.Uint16([]byte(peers[i+4 : i+6]))))
				fmt.Printf("%s.%s.%s.%s:%s\n",
					strconv.Itoa(int(peers[i])),
					strconv.Itoa(int(peers[i+1])),
					strconv.Itoa(int(peers[i+2])),
					strconv.Itoa(int(peers[i+3])),
					port)
			}
		}

	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
