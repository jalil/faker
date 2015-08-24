package internet

import (
	"encoding/json"
	"faker"
	"faker/name"
	"fmt"
	"math/rand"
	"strings"
)

const ALPHA = "qwertyuiopasdfghjkklzxcvbnm1234567890!@#$%^&*()+?>"

type InternetData struct {
	Internet struct {
		DomainSuffix []string `json:"domain_suffix"`
		FreeEmail    []string `json:"free_email"`
	} `json:"internet"`
}

func publicEmail() []string {
	internetdata := &InternetData{}

	data, err := faker.JsonData()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(data), &internetdata)
	if err != nil {
		panic(err)
	}

	return internetdata.Internet.DomainSuffix
}

func internetDomain() []string {
	internetdata := &InternetData{}

	data, err := faker.JsonData()
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(data), &internetdata)
	if err != nil {
		panic(err)
	}

	return internetdata.Internet.FreeEmail

}

func Url() string {
	return protocol() + strings.ToLower(name.FirstName()) + "." + publicEmail()[rand.Intn(len(publicEmail()))]
}

func Email() string {
	return UserName() + "@" + internetDomain()[rand.Intn(len(internetDomain()))]
}

func UserName() string {
	return strings.ToLower(name.FirstName()) + "." + strings.ToLower(name.LastName())
}

func Password() string {

	faker.Seeder()
	passwd := []byte(ALPHA)
	words := make([]string, 8)
	passwdlen := len(ALPHA)
	for i := 0; i < 9; i++ {
		words = append(words, string(passwd[rand.Intn(passwdlen)]))
	}
	return strings.Join(words, "")
}

func MacAddress() string {
	return "not yet implemented"
}

func protocol() string {
	protocol := "http://www."
	return protocol
}
