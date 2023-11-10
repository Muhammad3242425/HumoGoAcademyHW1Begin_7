package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64

	fmt.Println("Pi = 3.14")
	fmt.Print("R = ")
	
	fmt.Scan(&R)


	fmt.Println("L = ",2 *3.14 * R) 
	fmt.Println("S = ", 3.14* math.Pow(R,2))

}
