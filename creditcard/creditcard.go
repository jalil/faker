package creditcard

import (
	"faker"
	"math/rand"
)

func init() {
	faker.Seeder()
}

type CardData struct {
	CreditCard struct {
		CardType   []string `json:"card_types"`
		CardNumber []string `json:"card_numbers"`
	} `json:"creditcard"`
}

func cardType() []string {

	carddata := &CardData{}
	faker.UnmarshalData(carddata)
	return carddata.CreditCard.CardType

}

func cardNumber() []string {

	carddata := &CardData{}
	faker.UnmarshalData(carddata)
	return carddata.CreditCard.CardNumber

}
func Number() string {
	cardlen := len(cardNumber())
	return cardNumber()[rand.Intn(cardlen)]
}

//NEED TO FND A WAY TO RETURN TIME.DATE
//func Expiration() time {
//return time.Now()
//return time.AddDate(1, 2, 3)
//}

func CardType() string {
	cardlen := len(cardType())
	return cardType()[rand.Intn(cardlen)]
}
