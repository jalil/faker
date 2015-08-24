package app

import (
	"encoding/json"
	"faker"
	"faker/name"
	"fmt"
	"math/rand"
)

type AppData struct {
	App struct {
		Name []string
	}
}

func AppName() []string {
	appdata := &AppData{}

	data, err := faker.JsonData()

	err = json.Unmarshal([]byte(data), &appdata)
	if err != nil {
		fmt.Println(err)
	}

	return appdata.App.Name
}

func init() {
	faker.Seeder()
}

func Name() string {
	NameLen := len(AppName())
	return AppName()[rand.Intn(NameLen)]
}

//Version returns a float64  of the app Version
func Version() float64 {
	//NOT DONE.  two decimal places
	return rand.Float64() * 5
}

func AuthorName() string {
	return name.FullName()
}
