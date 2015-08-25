package address

import (
	"faker"
	"faker/name"
	"math/rand"
	"strconv"
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

func CityPrefix() []string {

	addressdata := &AddressData{}
	faker.UnmarshalData(addressdata)
	return addressdata.Address.CityPrefix
}

func stateData() []string {

	addressdata := &AddressData{}
	faker.UnmarshalData(addressdata)
	return addressdata.Address.State
}

func suffixData() []string {

	addressdata := &AddressData{}
	faker.UnmarshalData(addressdata)
	return addressdata.Address.StreetSuffix
}

func countryCode() []string {

	addressdata := &AddressData{}
	faker.UnmarshalData(addressdata)
	return addressdata.Address.CountryCode
}

func countryData() []string {

	addressdata := &AddressData{}
	faker.UnmarshalData(addressdata)
	return addressdata.Address.Country
}

func cityData() []string {

	addressdata := &AddressData{}
	faker.UnmarshalData(addressdata)
	return addressdata.Address.City
}

//This runs before any function call by default in GO.
//No need to seed rand in my functions!
func init() {
	faker.Seeder()
}

//Country return country name, panics if its cant read json data
func Country() string {
	return countryData()[rand.Intn(len(countryData()))]
}

//Country return country name, panics if its cant read json data
func CountryCode() string {
	return countryCode()[rand.Intn(len(countryCode()))]
}

//Country return country name, panics if its cant read json data
func Street() string {
	//suffix := suffixData()[rand.Intn(len(suffixData()))]
	suffix := suffixData()[rand.Intn(len(suffixData()))]
	return name.LastName() + " " + suffix
}

//Country return country name, panics if its cant read json data
func StreetAddress() string {
	return HouseNumber() + " " + Street()
}

//Country return country name, panics if its cant read json data
func City() string {
	return cityData()[rand.Intn(len(cityData()))]
}

//Country return country name, panics if its cant read json data
func State() string {
	return stateData()[rand.Intn(50)]
}

//Country return country name, panics if its cant read json data
func HouseNumber() string {
	streetnumber := rand.Intn(1000)
	return strconv.Itoa(streetnumber)
}

//func Postalcode() string {
//return faker.Postcode()[rand.Intn(50)]
//}

//Longitude return float64 longitude
func Longitude() float64 {
	return rand.Float64() * 180
}

//Latitude return float64 of latitude
func Latitude() float64 {
	return rand.Float64() * 180
}
