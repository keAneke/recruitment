// sorting2
package main

import "fmt"

func main() {
	var array = [6] int {1, 4, 5, 6, 8, 2}
	var highest = array[0]
	for i := 0; i <= 5; i++ {
		if array[i] > highest {
			highest = array[i]
		}
	}
	for i:= 1; i<n; i++ {		//sorting
		temp = array[i]
		j = i-1;
		while (temp<array[i] && j>=0){	//mengurutkan secara ascending
			array[j+1] = array[j]
			--j
		}
		array[j+1] = temp
	}
	for i :=0; i<n, i++ {
		fmt.Printf("%d\t", array[i])
	}
	
	for r := highest; r >= 1; r-- { //rows
		for c := 0; c <= 5; c++ { //cols
			if array[int i] >= r {
				fmt.Print("|")
			} else 
				fmt.Print(" ")
		}
	}
}
