package phone

import (
	"encoding/json"
	"faker"
	"fmt"
	"math/rand"
)

type PhoneData struct {
	Phone struct {
		AreaCode     []string `json:"area_code"`
		ExchangeCode []string `json:"exchange_code"`
	} `json:"phone_number"`
}

func phoneAreaCode() []string {
	phonedata := &PhoneData{}

	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &phonedata)
	if err != nil {
		fmt.Println(err)
	}

	return phonedata.Phone.AreaCode
}

func phoneExchangeCode() []string {
	phonedata := &PhoneData{}

	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &phonedata)
	if err != nil {
		fmt.Println(err)
	}

	return phonedata.Phone.AreaCode
}

func AreaCode() string {
	faker.Seeder()
	return phoneAreaCode()[rand.Intn(len(phoneAreaCode()))]
}

func exchangeCode() string {
	faker.Seeder()
	return phoneExchangeCode()[rand.Intn(len(phoneExchangeCode()))]

}

func Number() string {
	//faker.Seeder()
	//phonesuffix := rand.Intn(1000)
	//fmt.Println(AreaCode() + "-" + exchangeCode() + "-" + phonesuffix)
	return "Not yet"
}
