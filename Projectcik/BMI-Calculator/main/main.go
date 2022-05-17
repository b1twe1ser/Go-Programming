package main

import (
	"fmt"
	"strconv"
)

func calculateBMI(weight, height float64) (res float64) {
	// 👨🏼‍🏫 apply the formular BMI = weight / height^2
	res = weight / (height * height)
	return res
}

/*
  🚨 When fetching user input from the terminal,
  the input is of type string, hence the return type is a string
  (we will change that later)
*/
func getWeight() string {
	fmt.Println("Enter your weight: ")
	var weight string

	// &weight is a pointer to the value of weight
	// |-> The scanned value is written directly into the weight variable and stored there
	fmt.Scanln(&weight)
	return weight
}

/*
  🚨 When fetching user input from the terminal,
  the input is of type string, hence the return type is a string
  (we will change that later)
*/
func getHeight() string {
	fmt.Println("Enter your height: ")
	var height string
	fmt.Scanln(&height)
	return height
}

/*
  💡 Idea: We don't yet tell the user what their BMI means,
  as in is the BMI too high or too low ?
  Refer to the read me for more detail.
*/
func main() {

	fmt.Println("BMI Calculator")

	// Getting user weight and age (type string for now)
	var weightString = getWeight()
	var heightString = getHeight()

	// ⛔️ Conversion of type string to type float64
	weight, err := strconv.ParseFloat(weightString, 64)
	if err != nil {
		fmt.Println(err)
	}

	// ⛔️ Conversion of type string to type float64
	height, err := strconv.ParseFloat(heightString, 64)
	if err != nil {
		fmt.Println(err)
	}

	// calculating the BMI with weight and height
	var bmi = calculateBMI(weight, height)

	// printing the BMI to the terminal (formatted string, percision of 2 decimal places)
	fmt.Printf("Your BMI is %.2f \n", bmi)
}
