package app

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var n string
	name := app.Name()
	fmt.Println(name)

	//if n != "Twizzle" {
	//t.Error("Expected Twizzle, got", n)
	//}
	//}

	//func TestVersion(t *testing.T) {
	//var n float64
	//n = Verson()

	//if n != 1.0 {
	//t.Error("Expected 1.0, got", n)
	//}
}

func TestAuthorName(t *testing.T) {
	var n string
	n = Name()

	if n != "Zeus" {
		t.Error("Expected Zeus, got", n)
	}
}
