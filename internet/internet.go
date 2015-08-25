package internet

import (
	"faker"
	"faker/name"
	"math/rand"
	"strings"
)

type InternetData struct {
	Internet struct {
		DomainSuffix []string `json:"domain_suffix"`
		FreeEmail    []string `json:"free_email"`
	} `json:"internet"`
}

const ALPHA = "qwertyuiopasdfghjkklzxcvbnm1234567890!@#$%^&*()+?>"

func publicEmail() []string {

	internetdata := &InternetData{}
	faker.UnmarshalData(internetdata)
	return internetdata.Internet.DomainSuffix
}

func internetDomain() []string {

	internetdata := &InternetData{}
	faker.UnmarshalData(internetdata)
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
	//Need to make cleaner
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
