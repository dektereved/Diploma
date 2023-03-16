package service

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ReadFullFile reads a file and returns string or error
func ReadFullFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("файл не найден")
	}
	text := string(file)
	return text, nil
}

// MakeRequest sends Get request to a http adress and returns []byte or error
func MakeRequest(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, errors.New("ошибка соединения")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("ошибка. http-статус: resp.StatusCode")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ошибка чтения ответа")
	}
	return body, nil
}

// GetLines gets line out of the the text and returns []string
func GetLines(text string) []string {
	text = strings.ReplaceAll(text, "\"", "")
	lines := strings.Split(text, "\n")
	return lines
}

// GetWords gets a single word out of the text and returns []string
func GetWords(lines []string, i *int) []string {
	var words []string
	indx := *i
	words = strings.Split(lines[indx], ";")
	return words
}

// CheckCountryF compares the abbreviation of the country name in alpha with 2 to the data from the file and returns error
func CheckCountryF(words []string, listCountries map[int]string) error {
	for _, v := range listCountries {
		if words[0] == v {
			return nil
		}
	}
	return errors.New("некорректная страна")
}

// ConvertToInt converts string to int and returns int
func ConvertToInt(word string) int {
	converted, err := strconv.Atoi(word)
	if err != nil {
		log.Println(err)
	}
	return converted
}
