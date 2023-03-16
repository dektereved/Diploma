package voicecall

import (
	countrieslist "diploma/initial/countriesList"
	"diploma/service"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type VoiceCallData struct {
	Country       string
	Load          int
	ResponseTime  int
	Provider      string
	Stability     float32
	Purity        int
	Duration      int
	LastParameter int
}

func putToVoiceCallData(words []string) VoiceCallData {
	var value VoiceCallData
	for i, word := range words {
		if i == 0 {
			value.Country = word
		}
		if i == 1 {
			value.Load = service.ConvertToInt(word)
		}
		if i == 2 {
			value.ResponseTime = service.ConvertToInt(word)
		}
		if i == 3 {
			value.Provider = word
		}
		if i == 4 {
			value.Stability = ConvertToFloat(word)
		}
		if i == 5 {
			value.Purity = service.ConvertToInt(word)
		}
		if i == 6 {
			value.Duration = service.ConvertToInt(word)
		}
		if i == 7 {
			value.LastParameter = service.ConvertToInt(word)
		}
	}
	return value
}

func ConvertToFloat(word string) float32 {
	converted64, err := strconv.ParseFloat(word, 64)
	if err != nil {
		log.Println(err)
	}
	converted32 := float32(converted64)
	return converted32
}

func checkProvider(words []string) error {
	provider := words[3]
	if provider == "TransparentCalls" || provider == "E-Voice" || provider == "JustPhone" {
		return nil
	}
	return errors.New("некорректный провайдер")
}

func checkWords(words []string) error {
	if len(words) == 8 {
		return nil
	} else {
		return errors.New("некорректная длина строки")
	}
}

// VoiceCall gets data from the file and converts it into []VoiceCallData
func VoiceCall() {
	countrieslist := countrieslist.GetCountriesList()
	var VoiceCallStorrage []VoiceCallData
	text, err := service.ReadFullFile("test/voice.data")
	if err != nil {
		log.Println("Ошибка:", err)
		fmt.Printf("%+v\n", VoiceCallStorrage)
	}
	lines := service.GetLines(text)
	for i := 0; i < len(lines)-1; i++ {
		words := service.GetWords(lines, &i)
	innerLoop:
		for {
			err := checkWords(words)
			if err != nil {
				log.Println("Ошибка:", err)
				break innerLoop
			}
			err = service.CheckCountryF(words, countrieslist)
			if err != nil {
				log.Println("Ошибка:", err)
				break innerLoop
			}
			err = checkProvider(words)
			if err != nil {
				log.Println("Ошибка:", err)
				break innerLoop
			}
			value := putToVoiceCallData(words)
			VoiceCallStorrage = append(VoiceCallStorrage, value)
			break innerLoop
		}
	}
	for _, v := range VoiceCallStorrage {
		fmt.Printf("%+v\n", v)
	}
}
