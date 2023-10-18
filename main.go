package main

import "fmt"

func main() {
	for {
		var action uint8
		var number1 float64
		var number2 float64
		fmt.Println("$ Enter first number:")
		fmt.Scan(&number1)
		fmt.Println("$ Enter second number:")
		fmt.Scan(&number2)
		fmt.Println("$ 1) +, 2) -, 3) /, 4)*, 5)EXIT!")
		fmt.Scan(&action)
		if action == 1 {
			fmt.Println("$ Youre result is:\n", number1+number2)
		} else if action == 2 {
			fmt.Println("$ Youre result is:\n", number1-number2)
		} else if action == 3 {
			fmt.Println("$ Youre result is:\n", number1/number2)
			if number2 == 0 {
				fmt.Println("$ You can't / on zero!")
			}
		} else if action == 4 {
			fmt.Println("$ Youre result is:\n", number1*number2)
		} else if action == 5 {
			break
		}
	}
}
