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
