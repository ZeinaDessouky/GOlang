package myMath

import "fmt"

type calculator struct {
	num1 int
	num2 int
}

func GetCalc(a, b int) calculator {
	return calculator{
		num1: a,
		num2: b,
	}
}

func (calc calculator) Add() int {
	p := calc.num1 + calc.num2
	fmt.Println(p)
	return p
}

func (calc calculator) Subtract() int {
	p := calc.num1 - calc.num2
	fmt.Println(p)
	return p
}

func (calc calculator) Multiply() int {
	p := calc.num1 * calc.num2
	fmt.Println(p)
	return p
}

func (calc calculator) Divide() int {
	p := calc.num1 / calc.num2
	fmt.Println(p)
	return p
}
