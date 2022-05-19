package main

import "fmt"

func main() {

	// Task 1

	var middayTime = 13
	var timeNow = 12

	fmt.Println(middayTime == timeNow)

	// Task 2

	var numberOfFingersOnHand = 8
	var numberOfThumbs = 2

	fmt.Println(numberOfFingersOnHand == numberOfThumbs) // -> false
	fmt.Println(numberOfFingersOnHand != numberOfThumbs) // -> true
	fmt.Println(numberOfFingersOnHand < numberOfThumbs)  // -> false
	fmt.Println(numberOfFingersOnHand & numberOfThumbs)  // -> 0
	fmt.Println(numberOfFingersOnHand | numberOfThumbs)  // -> 10
	fmt.Println(numberOfFingersOnHand >= numberOfThumbs) // true
}
