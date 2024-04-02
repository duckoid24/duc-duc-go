package hello

import (
	"fmt"

	"check24.de/greeting"
)

func Hello(name string) {
	fmt.Println(greeting.Greeting(name))
}
