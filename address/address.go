package address

import (
	"faker"
	"faker/name"
	"math/rand"
	"strconv"
)

func init() {
	//This runs before any function call by default in GO.
	//No need to seed rand in my functions!
	faker.Seeder()
}
func Country() string {
	return faker.Country()[rand.Intn(len(faker.Country()))]
}

func CountryCode() string {
	return faker.CountryCode()[rand.Intn(len(faker.CountryCode()))]
}

func Street() string {
	suffix := faker.Suffix()[rand.Intn(len(faker.Suffix()))]
	return name.FirstName() + " " + suffix
}

func StreetAddress() string {
	return HouseNumber() + " " + Street()
}

func City() string {
	return faker.City()[rand.Intn(len(faker.City()))]
}

func State() string {
	return faker.State()[rand.Intn(50)]
}

func HouseNumber() string {
	streetnumber := rand.Intn(1000)
	return strconv.Itoa(streetnumber)
}

//func Postalcode() string {
//faker.Seeder()
//return rand.Floatn()
//return faker.Postcode()[rand.Intn(50)]
//}

func Longitude() float64 {
	return rand.Float64() * 180
}

func Latitude() float64 {
	return rand.Float64() * 180
}
