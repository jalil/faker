package phone

import (
	"faker"
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
	faker.UnmarshalData(phonedata)
	return phonedata.Phone.AreaCode
}

func phoneExchangeCode() []string {

	phonedata := &PhoneData{}
	faker.UnmarshalData(phonedata)
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
