package avatar

import (
	"faker"
	"faker/lorem"
	"math/rand"
)

func init() {
	faker.Seeder()
}
func Image() string {
	slug := lorem.LoremWord()[rand.Intn(loremWordsLen())]
	return "http://robohash.org/" + slug + ".png"
}

func ImageSize(size string) string {
	return Image() + "?" + string(size)
}

func loremWordsLen() int {
	return len(lorem.LoremWord())
}
