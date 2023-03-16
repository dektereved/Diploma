package mms

import (
	countrieslist "diploma/initial/countriesList"
	"diploma/service"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func Unmarshal(body []byte) ([]MMSData, error) {
	var MMSStorage []MMSData
	err := json.Unmarshal(body, &MMSStorage)
	if err != nil {
		return nil, errors.New("ошибка чтения данных")
	}
	return MMSStorage, nil
}

func getItem(MMSStorage []MMSData, i *int) []MMSData {
	var item []MMSData
	indx := *i
	item = append(item, MMSStorage[indx])
	return item
}

func checkCountry(item []MMSData, listCountries map[int]string) error {
	for _, v := range listCountries {
		if item[0].Country == v {
			return nil
		}
	}
	return errors.New("некорректная страна")
}

func checkProvider(item []MMSData) error {
	provider := item[0].Provider
	if provider == "Topolo" || provider == "Rond" || provider == "Kildy" {
		return nil
	}
	return errors.New("некорректный провайдер")
}

// Mms gets data from the http server and converts it into []MMSData
func Mms() {
	var MMSStorage []MMSData
	countrieslist := countrieslist.GetCountriesList()
	body, err := service.MakeRequest("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Println("Ошибка:", err)
		fmt.Printf("%+v\n", MMSStorage)
	}
	var item []MMSData
	MMSStorage, err = Unmarshal(body)
	if err != nil {
		log.Println("Ошибка:", err)
		fmt.Printf("%+v\n", MMSStorage)
	}
	for i := 0; i < len(MMSStorage); i++ {
		item = getItem(MMSStorage, &i)
		err = checkCountry(item, countrieslist)
		if err != nil {
			log.Println("Ошибка:", err)
			MMSStorage = append(MMSStorage[:i], MMSStorage[i+1:]...)
		}
		err = checkProvider(item)
		if err != nil {
			log.Println("Ошибка:", err)
			MMSStorage = append(MMSStorage[:i], MMSStorage[i+1:]...)
		}
	}
	for _, v := range MMSStorage {
		fmt.Printf("%+v\n", v)
	}
}
