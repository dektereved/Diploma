package billing

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func readFullFile() string {
	file, err := os.ReadFile("billing.data")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)
	return text
}

func convertToInts(bytes []byte) []int {
	var ints []int
	for _, v := range bytes {
		x, _ := strconv.Atoi(string(v))
		ints = append(ints, x)
	}
	return ints
}

func reverseInts(ints []int) []int {
	var result []int
	for i := len(ints) - 1; i >= 0; i-- {
		result = append(result, ints[i])
	}
	return result
}

func calculateInts(ints []int) uint8 {
	var numbersUint uint8
	for i, v := range ints {
		if v > 0 {
			numbersUint += uint8(math.Pow(2, float64(i)))
		}
	}
	return numbersUint
}

func putToBillingData(binaryNumber string) BillingData {
	var value BillingData
	var ints2 []int
	var result bool
	for i, v := range binaryNumber {
		x, _ := strconv.Atoi(string(v))
		if i == 0 && x != 0 {
			ints2 = append(ints2, 0, x)
		} else {
			ints2 = append(ints2, x)
		}
	}
	if ints2[0] > 0 {
		result = true
	} else {
		result = false
	}
	value.CreateCustomer = result
	if ints2[1] > 0 {
		result = true
	} else {
		result = false
	}
	value.Purchase = result
	if ints2[2] > 0 {
		result = true
	} else {
		result = false
	}
	value.Payout = result
	if ints2[3] > 0 {
		result = true
	} else {
		result = false
	}
	value.Recurring = result
	if ints2[4] > 0 {
		result = true
	} else {
		result = false
	}
	value.FraudControl = result
	if ints2[5] > 0 {
		result = true
	} else {
		result = false
	}
	value.CheckoutPage = result
	return value
}

// Billing gets data from the file and converts it into []BillingData
func Billing() {
	var BillingStorrage []BillingData
	var value BillingData
	var ints []int
	var numbersUint uint8
	text := readFullFile()
	bytes := []byte(text)
	for {
		if len(bytes) > 6 {
			bytes = bytes[:len(bytes)-1]
		} else {
			break
		}
	}
	ints = convertToInts(bytes)
	ints = reverseInts(ints)
	numbersUint = calculateInts(ints)
	binaryNumber := strconv.FormatInt(int64(numbersUint), 2)
	value = putToBillingData(binaryNumber)
	BillingStorrage = append(BillingStorrage, value)
	for _, v := range BillingStorrage {
		fmt.Printf("%+v\n", v)
	}

}
