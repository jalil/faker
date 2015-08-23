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

func Version() float64 {
	return rand.Float64() * 5
}

func AuthorName() string {
	return name.FullName()
}
