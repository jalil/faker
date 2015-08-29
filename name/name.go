package name

import (
	"faker"
	"math/rand"
)

type NameData struct {
	Name struct {
		FirstName  []string `json:"first_name"`
		LastName   []string `json:"last_name"`
		PrefixName []string `json:"prefix"`
	} `json:"name"`
}

func init() {
	faker.Seeder()
}

func FirstName() string {

	namedata := &NameData{}
	faker.UnmarshalData(namedata)
	fnameLen := len(namedata.Name.FirstName)
	return namedata.Name.FirstName[rand.Intn(fnameLen)]
}

func NamePrefix() string {

	namedata := &NameData{}
	faker.UnmarshalData(namedata)
	prefixLen := len(namedata.Name.PrefixName)
	return namedata.Name.PrefixName[rand.Intn(prefixLen)]
}

func LastName() string {

	namedata := &NameData{}
	faker.UnmarshalData(namedata)
	lnameLen := len(namedata.Name.LastName)
	return namedata.Name.LastName[rand.Intn(lnameLen)]
}

func FullName() string {
	return FirstName() + " " + LastName()
}
