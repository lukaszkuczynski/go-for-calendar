package go_for_calendar

import (
	"errors"
	"fmt"
	"time"
)

func say_hi(name string) string {
	var lineToWrite string = fmt.Sprintf("Welcome to the Kuczas World, %s", name)
	fmt.Println(lineToWrite)
	return lineToWrite
}

func give_me_5() (the_return_value int) {
	the_return_value = 5
	return
}

func main() {
	var name = "Kuczas"
	say_hi((name))
	fmt.Println("The time is", time.Now())
}

func aFunctionThatGetsWord() string {
	var a string
	a = "the_word"
	return a
}

// Playing with loops
func createNDigit(noOfDigits int) string {
	var result string
	// for i, value := range noOfDigits {
	// 	result += fmt.Sprint(i)

	// }
	for i := 1; i < noOfDigits+1; i++ {
		result += fmt.Sprint(i)
	}
	return result
}

// how to switch
func gimmeTextOfNumber(aNumber int) string {
	switch aNumber {
	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return "unknown"
	}
}

func countNoOfElementsInSlice(aSlice []int) int {
	return len(aSlice)
}

func getACapitalCity(country string) (string, error) {
	capitalCountryMap := make(map[string]string)
	capitalCountryMap["Poland"] = "Warsaw"
	capitalCountryMap["Germany"] = "Berlin"
	capitalCountryMap["France"] = "Paris"
	capital, ok := capitalCountryMap[country]
	if !ok {
		return "", errors.New("Country not found!")
	}
	return capital, nil
}
