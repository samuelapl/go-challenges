package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("Enter the first number: ")

	_, err := fmt.Scanf("%d", &a)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter the second number: ")

	_, err = fmt.Scanf("%d", &b)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return

	}

	result := Sum(a, b)

	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)

}

func Sum(a int, b int) int {
	return a + b
}
