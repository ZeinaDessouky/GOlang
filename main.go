package main

import "fmt"
import myMath "GOlang/myMath"

func main() {
	fmt.Println("Hello, World")

	calc := myMath.GetCalc(14, 7)

	calc.Add()
	calc.Subtract()
	calc.Multiply()
	calc.Divide()
}
