package main

import (
	"fmt"
	"testing"

	"TPseminarioGo.com/model"
	"github.com/stretchr/testify/assert"
)

func TestMode(t *testing.T) {
	var cases = []struct {
		Input  string // input string in order to be parsed
		Result bool   // parser result
		Type   string // the input type
		Value  string // the input value
		Length int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN0512345", true, "NN", "12345", 5},
	}
	for _, testData := range cases {
		r, err := model.GenerateStruct(testData.Input)
		if assert.Equal(t, err == nil, testData.Result) {
			if assert.Equal(t, r.Type, testData.Type) {
				fmt.Print("Pasó por el test")
			}
			if assert.Equal(t, r.Length, testData.Length) {
				fmt.Print("Pasó por el test")
			}
			if assert.Equal(t, r.Value, testData.Value) {
				fmt.Print("Pasó por el test")
			}
		}
	}

}
