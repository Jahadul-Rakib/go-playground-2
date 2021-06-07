package main

import (
	"errors"
	"log"
	"reflect"
)

func main() {

	number, err := addSomeNumber(12, 23.09)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(number)
	}
}

func addSomeNumber(num1 int8, numb2 float32) (int8, error) {
	if reflect.TypeOf(numb2).Kind() != reflect.Int8 {
		return 0, errors.New("type is not integer8")
	} else {
		return num1 + int8(numb2), nil
	}
}
