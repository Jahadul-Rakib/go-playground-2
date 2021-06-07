package main

import (
	"errors"
	"fmt"
	"log"
)

type Color struct {
	Red   bool
	Green bool
	Blue  bool
}
type Car struct {
	Model    string
	Price    float32
	CarColor Color
}

func (c Car) getCarInfo() (*Car, error) {
	if c.Price < 0 {
		return nil, errors.New("price is not actual")
	}
	return &c, nil
}

func main() {
	colorRed := Color{Red: true}

	carOne := Car{"BMW", -30000.12, colorRed}
	var mayCar, err = carOne.getCarInfo()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Print(mayCar)
	}
}
