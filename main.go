package main

import (
	"errors"
	"fmt"
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

func runStrangeLoopNTimes(n int) string {
	i := 1
	result := ""
	for i <= n {
		result += fmt.Sprint(i)
		i++
	}
	return result
}

func checkIfKuczek(aTextToCheck string) bool {
	if name := aTextToCheck; name == "Kuczek" {
		return true
	} else {
		return false
	}
}

type LukeDay struct {
	isBibleRead   bool
	areTeethClean bool
	workDayHours  int
}

func checkMyDay(day LukeDay) string {
	if !day.isBibleRead {
		return "not nice! Read Bible"
	} else if !day.areTeethClean {
		return "not nice! Go and wash your teeth!"
	} else {
		return "Good Boy! You worked " + fmt.Sprint(day.workDayHours) + " hours."
	}
}

func weekendGetter() []string {
	allDays := [7]string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	return allDays[5:6]
}

func makeASliceOfWows(size int) []string {
	aSlice := make([]string, size)
	for i := 0; i < size; i += 1 {
		aSlice[i] = "wow"
	}
	return aSlice
}

func returnStringMutateFunctionResult(mutator func(string) string, param string) string {
	return mutator(param)
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) calcInitials() string {
	return fmt.Sprint(string(p.firstName[0]), string(p.lastName[0]))
}

func (p *Person) setAge(age int) {
	p.age = age
}

type Identifiable interface {
	sayIntro() string
}

func (p Person) sayIntro() string {
	return "I am " + fmt.Sprint(p.firstName)
}

func checkIfStringBasedOnTypeAssertion(iSomething interface{}) bool {
	_, ok := iSomething.(string)
	return ok
}

func (p Person) String() string {
	return p.firstName
}

func printWithStringer(p Person) string {
	fmt.Sprint(p)
	return fmt.Sprint(p)
}

type PersonError struct {
	person Person
}

func (e *PersonError) Error() string {
	return "This Person has some error " + fmt.Sprint(e.person)
}

func (p *Person) getAgeValidated() (int, error) {
	if p.age < 0 {
		return 0, &PersonError{*p}
	}
	return p.age, nil
}
