package main

import (
	"fmt"
	"log"
)

func main() {

	var dataArray [10]string
	for index := range dataArray {
		dataArray[index] = "Hold On"
	}

	dataArr := [5]string{"A", "B", "C", "D", "E"}
	for i, data := range dataArr {
		log.Println(data, " Index Number ", i)
	}

	arr2D := [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	for _, dataCol := range arr2D {
		//log.Println(dataCol)
		for _, row := range dataCol {
			println(row)
		}
	}

	var a =[]int{1,2,3,4,4,5}
	b := append(a,32,52)
	for _, data := range b {
		log.Println(data)
	}

	makeSlice := make([]string, 3)
	for i, val := range makeSlice {
		makeSlice[i]= val+"Hell"
	}
	fmt.Print(makeSlice)

	var mp map[string]int = map[string]int{"Banana":23, "Watermelon": 200,"Cow":650}
	log.Println(mp)
	//or
	mapMake := make(map[string]string)

	mapMake["Rakib"]="SE"
	mapMake["Rakib1"]="SE1"
	mapMake["Rakib2"]="SE2"
	mapMake["Rakib3"]="SE3"

	delete(mapMake, "Rakib1")
	log.Println(mapMake)

	//checking key is present in map or not
	//If key is in map, ok is true. If not, ok is false

	value, ok :=mapMake["Rakib2"]
	if ok == true {
		log.Println("value is exist, ",value)
	} else {
		log.Println("value is not exist in the map, ", value)
	}

}
