package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func GenerateStruct(str string) (Result, error) {
	r := Result{
		Type:   "",
		Length: 0,
		Value:  "",
	}

	resultLength := len([]rune(str))

	if resultLength < 5 {
		return r, errors.New("la cadena no cumple con los requerimientos mínimos")
	}

	r.Type = str[:2]
	z, err := strconv.Atoi(str[2:4])
	if err != nil {
		return r, errors.New("insuficiencia de datos en la sección: Longitud")
	}
	r.Value = str[4:]
	r.Length = z

	if (resultLength - 4) != r.Length {
		return r, errors.New("el largo no coincide")
	}

	if r.Type != "TX" && r.Type != "NN" {
		return r, errors.New("no es de un tipo válido")
	}

	if r.Type == "NN" {
		if _, err := strconv.Atoi(r.Value); err != nil {
			return r, errors.New("no es de tipo número")
		}
	} else if r.Type == "TX" {
		if _, err := strconv.Atoi(r.Value); err != nil {
			return r, errors.New("no es de tipo texto")
		}
	}

	// for i := r.Value[4]; i < byte(s[len(s)-1]); i++ {
	// 	if (r.Value[i] >= 'a') && (r.Value[i] <= 'z') {
	// 		isChar = true
	// 		if r.Type == "TX" && !isChar {
	// 			return r, errors.New("type and value mismatch")
	// 		}
	// 	}
	// }

	return r, nil
}
