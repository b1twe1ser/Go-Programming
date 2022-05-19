package main

/*
	💡 Just ignore the red lines saying the variables are unused,
	it's not a problem unless you want to run the code. In that case,
	maybe try to use them or just comment them out
*/

func main() {

	// Task 1
	var myAge uint8 = 25
	var myAgeDoubled int64 = int64(myAge) * 2
	var myAgeDoubledDivided float32 = float32(myAgeDoubled) / 3.0
	var myMotto string = "I like programming 🔥"
	var myTruth string = myMotto + ", it's so easy 💗"

	// Task 2
	var myAge uint8 = 25
	var numberOfDoenerICanEatInADay uint8 = 3

	// We can't change the size of this variable as the value is too large for any smaller integer type
	// -> The result of conversion would not be the expected value
	var ageOfTheWorld int = 4.543e9

	// We can't decrease the size of a boolean
	var programmingIsFun bool = true
	var myAgeInTwoYares uint8 = myAge + 2

	// Since strings are essentially arrays or "A list of bytes" I suppose this conversion counts
	var firstLetterOfAlphabet byte = 'a'

	// Task 3

	// 🪲 1
	// func main(){} instead of func man(){}

	// 🪲 2
	// name := "Nico" instead of name = "Nico"

	// 🪲 3
	// fmt.Pritnln(name + age) won't work because we are trying to add a string and a integer together
	// -> Not compatible -> program will crash
}
