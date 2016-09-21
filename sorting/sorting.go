// sorting
package main

import "fmt"

func main() {
	var array = [6]int{1, 4, 5, 6, 8, 2}
	//fmt.Print(array)
	var highest = array[0]
	for i := 0; i <= 5; i++ {
		if array[i] > highest {
			highest = array[i]
		}
	}
	for r := highest; r >= 1; r-- { //rows
		for c := 0; c <= 5; c++ { //cols
			if array >= r {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
	}
}
