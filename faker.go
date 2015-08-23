package faker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type AddressData struct {
	Address struct {
		CityPrefix   []string `json:"city_prefix"`
		City         []string `json:"city_suffix"`
		Country      []string `json:"country"`
		CountryCode  []string `json:"country_code"`
		State        []string `json:"state"`
		StateCode    []string `json:"state_abbr"`
		StreetSuffix []string `json:"street_suffix"`
	} `json:"address"`
}

type InternetData struct {
	Internet struct {
		DomainSuffix []string `json:"domain_suffix"`
		FreeEmail    []string `json:"free_email"`
	} `json:"internet"`
}

type CardData struct {
	CreditCard struct {
		CardType   []string `json:"card_types"`
		CardNumber []string `json:"card_numbers"`
	} `json:"creditcard"`
}

type PhoneData struct {
	Phone struct {
		AreaCode     []string `json:"area_code"`
		ExchangeCode []string `json:"exchange_code"`
	} `json:"phone_number"`
}

func JsonData() ([]byte, error) {

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/en-US.json")

	return data, err
}

func Seeder() {
	rand.Seed(time.Now().UnixNano())
}

func RandNumber(data int) int {
	return rand.Intn(data)
}

func PhoneAreaCode() []string {
	nameD := &PhoneData{}

	//data, _ := ioutil.ReadFile("phony.json")
	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/locale/en-US.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		fmt.Println(err)
	}

	return nameD.Phone.AreaCode
}
func InternetFreeEmail() []string {
	nameD := &InternetData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Internet.FreeEmail

}

func CardType() []string {
	nameD := &CardData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.CreditCard.CardType

}

func CardNumber() []string {
	nameD := &CardData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.CreditCard.CardNumber

}
func InternetDomain() []string {
	nameD := &InternetData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Internet.DomainSuffix

}

func CityPrefix() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.CityPrefix
}

func CountryCode() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.CountryCode
}

func Country() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.Country
}

func City() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.City
}
func State() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.State
}
func StateCode() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.StateCode
}
func Suffix() []string {
	nameD := &AddressData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.StreetSuffix
}
func LoadJson() []string {
	nameD := &AddressData{}

	data, _ := ioutil.ReadFile("phony.json")

	err := json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Address.CountryCode
}

func LoadConfig(path string) (map[string]interface{}, error) {
	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)
	return m, err
}
