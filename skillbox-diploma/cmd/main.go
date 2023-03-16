package main

import (
	"diploma/initial/billing"
	"diploma/initial/email"
	"diploma/initial/incidents"
	"diploma/initial/mms"
	"diploma/initial/sms"
	"diploma/initial/support"
	"diploma/initial/voicecall"
	"fmt"
)

func main() {
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ SMS:")
	sms.Sms()
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ MMS:")
	mms.Mms()
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ VOICECALL:")
	voicecall.VoiceCall()
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ EMAIL:")
	email.Email()
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ BILLING:")
	billing.Billing()
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ SUPPORT:")
	support.Support()
	fmt.Println("РЕЗУЛЬТАТ РАБОТЫ ФУНКЦИИ INCIDENTS И ЗАПУСК СЕРВЕРА:")
	incidents.Incidents()

}
