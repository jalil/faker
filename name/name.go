package name

import (
	"encoding/json"
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

func NameFirst() []string {
	nameD := &NameData{}

	//read in json data
	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Name.FirstName
}

func NamePrefix() []string {
	nameD := &NameData{}

	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Name.PrefixName
}

func NameLast() []string {
	nameD := &NameData{}

	data, err := faker.JsonData()
	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Name.LastName
}

func init() {
	faker.Seeder()
}

func FullName() string {
	return FirstName() + " " + LastName()
}

func FirstName() string {
	return NameFirst()[rand.Intn(len(NameFirst()))]
}

func Prefix() string {
	return NamePrefix()[rand.Intn(len(NamePrefix()))]
}

func LastName() string {
	return NameLast()[rand.Intn(len(NameLast()))]
}
