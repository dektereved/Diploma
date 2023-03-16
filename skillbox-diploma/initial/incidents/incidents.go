package incidents

import (
	"diploma/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func unmarshal(body []byte) ([]IncidentData, error) {
	var IncidentStorrage []IncidentData
	err := json.Unmarshal(body, &IncidentStorrage)
	if err != nil {
		return nil, err
	}
	return IncidentStorrage, nil
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Incidents gets data from the http server and converts it into []IncidentData
func Incidents() {
	var IncidentStorrage []IncidentData
	var empty IncidentData
	body, err := service.MakeRequest("http://127.0.0.1:8383/accendent")
	if err != nil {
		fmt.Println("Работа прервана:", err)
		IncidentStorrage = append(IncidentStorrage, empty)
	}
	IncidentStorrage, err = unmarshal(body)
	if err != nil {
		fmt.Println("Работа прервана:", err)
		IncidentStorrage = append(IncidentStorrage, empty)
	}
	fmt.Println(IncidentStorrage)
	for _, v := range IncidentStorrage {
		fmt.Printf("%+v\n", v)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", handleConnection)
	log.Fatal(http.ListenAndServe(":8282", router))

}
