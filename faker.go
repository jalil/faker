//Base package
package faker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

//JsonData reads in json file
func JsonData() ([]byte, error) {
	data, err := ioutil.ReadFile("en-US.json")
	if err != nil {
		//maybe use panic()?
		fmt.Println(err)
	}

	return data, err
}

//Seed sets a random seed base on time
func Seeder() {
	rand.Seed(time.Now().UnixNano())
}

//Unmarshal any input type passed in as parameter
func UnmarshalData(input interface{}) {
	data, err := JsonData()

	if err != nil {
		//maybe use panic()?
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(data), &input)
	//maybe use panic()?
	if err != nil {
		fmt.Println(err)
	}
}
