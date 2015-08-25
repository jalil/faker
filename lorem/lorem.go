package lorem

import (
	"faker"
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
	faker.UnmarshalData(loremdata)
	return loremdata.Lorem.Words

}

func loremSentence() []string {
	//NOT DONE
	return LoremWords()
}

//Runs before main
func init() {
	faker.Seeder()
}

func Word() string {
	loremLen := len(LoremWords())
	return LoremWords()[rand.Intn(loremLen)]
}

func Text() []string {
	//NOT DONE!
	return loremSentence()
}

func Character() string {
	return string(ALPHABETS[rand.Intn(54)])
}
