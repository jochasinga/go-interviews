package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

type Element struct {
	Id string
}

var getUniqueId = func() (func(Element) string) {
	nextGeneratedId := 435
	return func(elm Element) string {
		if elm.Id == "" {
			elm.Id = "generated-uid-" + strconv.Itoa(nextGeneratedId)
			nextGeneratedId = rand.Intn(1000)
		}
		return elm.Id
	}
}

func main() {
	e := Element{}
	genId := getUniqueId()
	fmt.Println(genId(e))

}
