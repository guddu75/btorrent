package main

import (
	"strconv"
	"unicode"
)

func DecodeString(bencodedString string, idx int) (string, int, error) {

	var firstColonIndex int

	for i := idx; i < len(bencodedString); i++ {
		if bencodedString[i] == ':' {
			firstColonIndex = i
			break
		}
	}

	lengthStr := bencodedString[idx:firstColonIndex]

	// fmt.Println(lengthStr)

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", 0, err
	}
	// fmt.Println(bencodedString[firstColonIndex+1:firstColonIndex+1+length], length+len(lengthStr)+1)
	return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], length + len(lengthStr) + 1, nil

}

func DecodeInt(bencodedString string, idx int) (int, int, error) {
	var num string
	neg := false
	for i := idx + 1; i < len(bencodedString); i++ {
		if unicode.IsDigit(rune(bencodedString[i])) {
			num = num + string(bencodedString[i])
		} else if rune(bencodedString[i]) == '-' {
			neg = true
		} else {
			break
		}
	}
	a, err := strconv.Atoi(num)
	if err != nil {
		return 0, 0, err
	}
	if neg {
		a = a * -1
	}
	return a, len(num) + 2, nil
}

func DecodeList(bencodedString string, idx int) ([]interface{}, int, error) {
	slice := make([]interface{}, 0)
	var i int
	for i = idx + 1; i < len(bencodedString); {
		if unicode.IsDigit(rune(bencodedString[i])) {
			decodedString, length, err := DecodeString(bencodedString, i)
			if err != nil {
				return nil, 0, err
			}
			slice = append(slice, decodedString)
			i += length
		} else if rune(bencodedString[i]) == 'i' {
			decodedINT, length, err := DecodeInt(bencodedString, i)
			if err != nil {
				return nil, 0, err
			}
			slice = append(slice, decodedINT)
			i += length
		} else if rune(bencodedString[i]) == 'l' {
			// fmt.Println("inside list")
			decodedList, length, err := DecodeList(bencodedString, i)
			if err != nil {
				return nil, 0, err
			}
			slice = append(slice, decodedList)
			i += length
		} else {
			break
		}
		// fmt.Println(slice...)
	}

	return slice, i - idx + 1, nil

}
