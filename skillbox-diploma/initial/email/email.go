package email

import (
	countrieslist "diploma/initial/countriesList"
	"diploma/service"
	"errors"
	"fmt"
	"log"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func putToEmailData(words []string) EmailData {
	var value EmailData
	for i, word := range words {
		if i == 0 {
			value.Country = word
		}
		if i == 1 {
			value.Provider = word
		}
		if i == 2 {
			value.DeliveryTime = service.ConvertToInt(word)
		}
	}
	return value
}

func checkProvider(words []string) error {
	provider := words[1]
	if provider == "Gmail" || provider == "Yahoo" || provider == "Hotmail" || provider == "MSN" || provider == "Orange" || provider == "Comcast" || provider == "AOL" || provider == "Live" || provider == "RediffMail" || provider == "GMX" || provider == "Protonmail" || provider == "Yandex" || provider == "Mail.ru" {
		return nil
	}
	return errors.New("некорректный провайдер")
}

func checkWords(words []string) error {
	if len(words) == 3 {
		return nil
	} else {
		return errors.New("некорректная длина строки")
	}
}

// Email gets data from the file and converts it into []EmailData
func Email() {
	countrieslist := countrieslist.GetCountriesList()
	var EmailStorrage []EmailData
	text, err := service.ReadFullFile("test/email.data")
	if err != nil {
		log.Println("Ошибка:", err)
		fmt.Printf("%+v\n", EmailStorrage)
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
			value := putToEmailData(words)
			EmailStorrage = append(EmailStorrage, value)
			break innerLoop
		}
	}
	for _, v := range EmailStorrage {
		fmt.Printf("%+v\n", v)

	}
}
