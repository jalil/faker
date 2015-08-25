package app

import (
	"faker"
	"faker/name"
	"math/rand"
)

type AppData struct {
	App struct {
		Name []string
	}
}

func init() {
	faker.Seeder()
}

func AppName() []string {

	appdata := &AppData{}
	faker.UnmarshalData(appdata)
	return appdata.App.Name
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
