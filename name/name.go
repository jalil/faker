package name

import (
	"encoding/json"
	"faker"
	"fmt"
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
	namedata := &NameData{}

	//read in json data
	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &namedata)
	if err != nil {
		fmt.Println(err)
	}

	return namedata.Name.FirstName
}

func NamePrefix() []string {
	namedata := &NameData{}

	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &namedata)
	if err != nil {
		panic(err)
	}

	return namedata.Name.PrefixName
}

func NameLast() []string {
	namedata := &NameData{}

	data, err := faker.JsonData()
	err = json.Unmarshal([]byte(data), &namedata)
	if err != nil {
		panic(err)
	}

	return namedata.Name.LastName
}

func init() {
	faker.Seeder()
}

func FullName() string {
	return FirstName() + " " + LastName()
}

func FirstName() string {
	return NameFirst()[0]
	//return NameFirst()[rand.Intn(len(NameFirst()))]
}

func Prefix() string {
	return NamePrefix()[rand.Intn(len(NamePrefix()))]
}

func LastName() string {
	return NameLast()[rand.Intn(len(NameLast()))]
}
