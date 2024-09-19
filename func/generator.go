package web

import "fmt"

func GenerateAscii(str string, fileArray []string) (string, error) {
	result := ""
	for line := 0; line <= 7; line++ {
		for _, char := range str {
			location, error := charLocator(rune(char))
			if error != nil {
				return "", error
			}
			result += fileArray[location+line]
		}
		result += "\n"
	}
	return result, nil
}

func charLocator(char rune) (int, error) {
	if char >= 32 && char <= 126 {
		location := 1 + 9*(char-32)
		return int(location), nil
	} else {
		return -1, fmt.Errorf("error Bad Request: Character %c is not in the range of 32-126", char)
	}
}
