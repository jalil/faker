package app

import (
	"faker/app"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var n string
	name := app.Name()
	fmt.Println(name)

}

func TestAuthorName(t *testing.T) {
	var n string
	n = Name()

	if n != "Zeus" {
		t.Error("Expected Zeus, got", n)
	}
}
