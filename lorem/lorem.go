package lorem

import (
	"encoding/json"
	"faker"
	"fmt"
	"math/rand"
)

const ALPHABETS = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

type LoremData struct {
	Lorem struct {
		Words []string `json:"words"`
	} `json:"lorem"`
}

func LoremWords() []string {
	loremdata := &LoremData{}

	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &loremdata)
	if err != nil {
		fmt.Println(err)
	}

	return loremdata.Lorem.Words

}

func loremSentence() []string {
	//NOT DONE
	return loremWords()
}

//Runs before main
func init() {
	faker.Seeder()
}

func Word() string {
	loremLen := len(loremWords())
	return loremWords()[rand.Intn(loremLen)]
}

func Text() []string {
	//NOT DONE!
	return loremSentence()
}

func Character() string {
	return string(ALPHABETS[rand.Intn(54)])
}
