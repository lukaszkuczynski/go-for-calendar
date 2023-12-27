package main

import (
	"fmt"
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

func TestStrangeLoop(t *testing.T) {
	expected := "12345"
	actual := runStrangeLoopNTimes(5)
	assert.Equal(t, expected, actual)
}

func TestShortIfAssignment(t *testing.T) {
	assert.Equal(t, true, checkIfKuczek("Kuczek"))
}

func TestMyDayStruct(t *testing.T) {
	lukeDay := LukeDay{false, false, 8}
	actual_message := checkMyDay(lukeDay)
	assert.Contains(t, actual_message, "not nice")
}

func TestMyDayStructWithLiteral(t *testing.T) {
	lukeDay := LukeDay{workDayHours: 6, isBibleRead: true, areTeethClean: true}
	actual_message := checkMyDay(lukeDay)
	assert.Contains(t, actual_message, "6 hours")
}

func TestWeekendGetter(t *testing.T) {
	weekendDays := weekendGetter()
	assert.Contains(t, weekendDays, "sat")
}

func TestMakeIt(t *testing.T) {
	wowsSlice := makeASliceOfWows(5)
	assert.Contains(t, wowsSlice, "wow")
	assert.Equal(t, len(wowsSlice), 5)
	assert.Equal(t, cap(wowsSlice), 5)
}

func TestMakeItAppended(t *testing.T) {
	// this is to show I can extend the original slice size
	wowsSlice := makeASliceOfWows(5)
	wowsSlice = append(wowsSlice, "wow")
	assert.Equal(t, len(wowsSlice), 6)
}

func TestWowsWithRange(t *testing.T) {
	wowsSlice := makeASliceOfWows(10)
	expectedI := 0
	for i, wow := range wowsSlice {
		assert.Equal(t, expectedI, i)
		assert.Equal(t, wow, "wow")
		expectedI += 1
	}
}

func TestMutatorIsLengthDescriptor(t *testing.T) {
	var lenghtMutator = func(txt string) string {
		return "Length is " + fmt.Sprint(len(txt))
	}
	aResult := returnStringMutateFunctionResult(lenghtMutator, "text")
	assert.Equal(t, "Length is 4", aResult)
}
