package main

import (
	"fmt"
	"testing"

	"TPseminarioGo.com/model"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var cases = []struct {
		Input  string // input string in order to be parsed
		Result bool   // parser result
		Type   string // the input type
		Value  string // the input value
		Length int    // value length
	}{
		{"TX04ABC", false, "TX", "ABC", 4},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"NN0884571236", true, "NN", "84571236", 8},
		{"TX06ABCDE", false, "TX", "ABCDE", 6},
		{"NN0512345", true, "NN", "12345", 5},
	}
	for _, testData := range cases {
		r, err := model.GenerateStruct(testData.Input)
		assert.Equal(t, err == nil, testData.Result)
		if err != nil {
			// if assert.Equal(t, err == nil, testData.Result) {
			// if assert.Equal(t, r.Type, testData.Type) {
			// 	fmt.Print("Pasó por el test")
			// }
			// if assert.Equal(t, r.Length, testData.Length) {
			// 	fmt.Print("Pasó por el test")
			// }
			// if assert.Equal(t, r.Value, testData.Value) {
			// 	fmt.Print("Pasó por el test")
			// }
			if (assert.Equal(t, r.Type, testData.Type)) &&
				(assert.Equal(t, r.Length, testData.Length)) &&
				(assert.Equal(t, r.Value, testData.Value)) {
				fmt.Print("Pasó por el test")
			}

		}
	}
}

//}

// func TestFoo(t *testing.T) {
// 	t.Logf("Testing Foo")
// }
