package main

import (
	"fmt"

	"TPseminarioGo.com/model"
)

func main() {
	r, e := model.GenerateStruct("NN0512345")
	if e == nil {
		fmt.Println(r)
	} else {
		fmt.Println(e)
	}
}
