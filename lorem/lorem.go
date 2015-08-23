package lorem

import (
	"encoding/json"
	"faker"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type LoremData struct {
	Lorem struct {
		Words []string `json:"words"`
	} `json:"lorem"`
}

func LoremWords() []string {
	loremdata := &LoremData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &loremdata)
	if err != nil {
		fmt.Println(err)
	}

	return loremdata.Lorem.Words

}

func LoremSentence() []string {
	nameD := &LoremData{}

	data, err := ioutil.ReadFile("/Users/jalil/go/src/faker/phony.json")

	err = json.Unmarshal([]byte(data), &nameD)
	if err != nil {
		panic(err)
	}

	return nameD.Lorem.Words

}

const ALPHABETS = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func init() {
	faker.Seeder()
}

func Word() string {
	loremLen := len(LoremWords())
	return LoremWords()[rand.Intn(loremLen)]
}

func Text() []string {
	return LoremSentence()
}

func Character() string {
	return string(ALPHABETS[rand.Intn(54)])
}
