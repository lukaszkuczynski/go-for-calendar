package go_for_calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloForAName(t *testing.T) {

	var expected string = "Welcome to the Kuczas World, Ziom"
	var actual string = say_hi("Ziom")

	assert.Equal(t, expected, actual)

}

func TestFor5(t *testing.T) {

	var expected int = 5
	var actual int = give_me_5()

	assert.Equal(t, expected, actual)

}