//Base package
package faker

import (
	"io/ioutil"
	"math/rand"
	"time"
)

//JsonData reads in json file
func JsonData() ([]byte, error) {
	data, err := ioutil.ReadFile("en-US.json")

	return data, err
}

//Seed sets a random seed base on time
func Seeder() {
	rand.Seed(time.Now().UnixNano())
}
