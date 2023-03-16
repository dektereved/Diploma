package sms

import (
	countrieslist "diploma/initial/countriesList"
	"diploma/service"
	"errors"
	"fmt"
	"log"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func putToSMSSData(words []string) SMSData {
	var value SMSData
	for i, word := range words {
		if i == 0 {
			value.Country = word
		}
		if i == 1 {
			value.Bandwidth = word
		}
		if i == 2 {
			value.ResponseTime = word
		}
		if i == 3 {
			value.Provider = word
		}
	}
	return value
}

func checkProvider(words []string) error {
	provider := words[3]
	if provider == "Topolo" || provider == "Rond" || provider == "Kildy" {
		return nil
	}
	return errors.New("некорректный провайдер")
}

func checkWords(words []string) error {
	if len(words) == 4 {
		return nil
	} else {
		return errors.New("некорректная длина строки")
	}
}

// Sms gets data from the file and converts it into []SMSData
func Sms() {
	countrieslist := countrieslist.GetCountriesList()
	var SMSstorage []SMSData
	text, err := service.ReadFullFile("test/SMS.data")
	if err != nil {
		log.Println("Ошибка:", err)
		fmt.Printf("%+v\n", SMSstorage)
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
			value := putToSMSSData(words)
			SMSstorage = append(SMSstorage, value)
			break innerLoop
		}
	}
	for _, v := range SMSstorage {
		fmt.Printf("%+v\n", v)
	}
}
