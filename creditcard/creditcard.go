package main

import (
	"faker"
	"fmt"
	"math/rand"
)

func init() {
	faker.Seeder()
}
func Number() string {
	cardlen := len(faker.CardNumber())
	return faker.CardNumber()[rand.Intn(cardlen)]
}

//NEED TO FND A WAY TO RETURN TIME.DATE
//func Expiration() time {
//return time.Now()
//return time.AddDate(1, 2, 3)
//}

func CardType() string {
	cardlen := len(faker.CardType())
	return faker.CardType()[rand.Intn(cardlen)]
}

func main() {
	fmt.Println(Number())
	fmt.Println(CardType())
}
