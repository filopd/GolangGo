package main

import "fmt"

func main() {
	fmt.Print("1. Print():", "Don't", "Have", "Spaces.")
	fmt.Print("2. Print():", "NotGoingToNextLine.")
	fmt.Println("3. Println():", "Have", "Spaces.")
	fmt.Println("4. Println():", "GoingToNextLine.")
	thisString := "Example"
	thisInt := 3
	thisFloat := 3.423
	fmt.Printf("5. Printf(): String %v, Number %d, Float %.2f", thisString, thisInt, thisFloat)
	thisFloat = 5.83
	varRunTime := fmt.Sprint("   6. Sprint():", "NoSpace & Float is", thisFloat)
	fmt.Println(varRunTime)
	thisFloat = 5.55
	varRunTime = fmt.Sprintln("7. Sprintln():", "With Space & new line. Float is", thisFloat)
	fmt.Print(varRunTime)
	thisFloat = 5.55421
	varRunTime = fmt.Sprintf("8. Sprintf(): Float is %.3f", thisFloat)
	fmt.Print(varRunTime)
	fmt.Println("\n\n9-1. Enter a Number:")
	fmt.Scan(&varRunTime)
	fmt.Println("9-2. You have entered:", varRunTime)
}
