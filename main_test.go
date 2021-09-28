package main

import (
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
		_, err := model.GenerateStruct(testData.Input)
		assert.Equal(t, err == nil, testData.Result)
	}
}
