package main

import "log"

func main() {
	x := 1
	for x <= 10 {
		log.Println(x, "printed")
		x++
	}

	for i := 0; i < 15; i++ {
		if i == 13 {
			log.Println("enter into break block")
			break
		} else if i == 5 {
			log.Println("now in 5")
			continue
		}
		log.Println(i+1, " time round")
	}

	a := 12

	switch a {
	case 1, 2, 34, 54:
		log.Println("Not Match")
	case 12:
		log.Println("Match")
	default:
		log.Println("from default")
	}

	b := 11

	switch {
	case b < 1:
		log.Println("less")
	case 12 > b:
		log.Println("high")
	default:
		log.Println("from default not match")
	}

}
