package utils

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func Notify(title string, body string) {
	fmt.Printf("%v: %v\n", title, body)
	if err := beeep.Notify(title, body, ""); err != nil {
		panic(err)
	}
}
