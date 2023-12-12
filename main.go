package go_for_calendar

import (
	"fmt"
	"time"
)

func say_hi(name string) string {
	var lineToWrite string = fmt.Sprintf("Welcome to the Kuczas World, %s", name)
	fmt.Println(lineToWrite)
	return lineToWrite
}

func main() {
	var name = "Kuczas"
	say_hi((name))
	fmt.Println("The time is", time.Now())
}
