package web

import (
	"errors"
	"net/http"
	"strings"
)

func Validation(wordToFind string, banner string, w http.ResponseWriter) (string, string, error) {
	file := "standard.txt" //default file
	//var file string
	if banner != "" {
		switch strings.ToLower(banner) {
		case "standard":
			file = "standard.txt"
		case "shadow":
			file = "shadow.txt"
		case "thinkertoy":
			file = "thinkertoy.txt"
		default:
			file = "standard.txt"

		}
	}

	if wordToFind == "" {
		return "", "", errors.New("error bad request: word empty")
	} else if len(wordToFind) > 300 {
		//  return empty string and error
		return "", "", errors.New("error bad request: word is too long")
	}

	return wordToFind, file, nil
}
