package go_for_calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloForAName(t *testing.T) {
	expected := "Welcome to the Kuczas World, Ziom"
	var actual string = say_hi("Ziom")

	assert.Equal(t, expected, actual)

}

func TestFor5(t *testing.T) {

	var expected rune = 5
	var actual rune = rune(give_me_5())

	assert.Equal(t, expected, actual)

}

func TestForAFunctionThatGetsWord(t *testing.T) {
	var actual = aFunctionThatGetsWord()
	assert.Equal(t, actual, "the_word")
}

func TestForNDigitText(t *testing.T) {
	var actual = createNDigit(4)
	expected := "1234"
	assert.Equal(t, expected, actual)
}

func TestForNumberSwitch(t *testing.T) {
	actual := gimmeTextOfNumber(2)
	expected := "two"
	assert.Equal(t, actual, expected)
}

func TestForSliceSizeCounting(t *testing.T) {
	var someSlice = make([]int, 5)
	// someArray[3] = 1
	// someArray[2] = 2
	actual := countNoOfElementsInSlice(someSlice)
	expected := 5
	assert.Equal(t, expected, actual)
}

func TestGetRange(t *testing.T) {
	var sliceDiv10 = [5]int{10, 20, 30, 40}
	assert.Contains(t, sliceDiv10, 10)
	assert.NotContains(t, sliceDiv10[1:2], 30)
}

func TestMapReturnsCorrectStuff(t *testing.T) {
	expected := "Paris"
	actual, _ := getACapitalCity("France")
	assert.Equal(t, expected, actual)
	actual2, errors2 := getACapitalCity("ThatCountryDoesntExist")
	assert.Equal(t, "", actual2)
	assert.NotNil(t, errors2)
}
