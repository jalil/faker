package main

import (
	"faker"
	"fmt"
	"math/rand"
)

func AreaCode() string {
	faker.Seeder()
	return faker.PhoneAreaCode()[rand.Intn(10)]
}

func exchangeCode() string {
	faker.Seeder()
	return faker.PhoneAreaCode()[rand.Intn(10)]

}

func Number() string {
	faker.Seeder()
	return AreaCode() + "-" + exchangeCode() + "-" + string(rand.Intn(1000))
}
func main() {
	fmt.Println(Number())
}
