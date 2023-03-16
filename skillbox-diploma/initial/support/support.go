package support

import (
	"diploma/service"
	"encoding/json"
	"fmt"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func Unmarshal(body []byte) ([]SupportData, error) {
	var SupportStorrage []SupportData
	err := json.Unmarshal(body, &SupportStorrage)
	if err != nil {
		return nil, err
	}
	return SupportStorrage, nil
}

// Support gets data from the http server and converts it into []SupportData
func Support() {
	var SupportStorrage []SupportData
	var empty SupportData
	body, err := service.MakeRequest("http://127.0.0.1:8383/support")
	if err != nil {
		fmt.Println("Работа прервана:", err)
		SupportStorrage = append(SupportStorrage, empty)
	}
	SupportStorrage, err = Unmarshal(body)
	if err != nil {
		fmt.Println("Работа прервана:", err)
		SupportStorrage = append(SupportStorrage, empty)
	}
	for _, v := range SupportStorrage {
		fmt.Printf("%+v\n", v)
	}
}
