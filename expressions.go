package main

import "fmt"

type MathExpr = string

const (
	Addition = MathExpr("addition")
	Subtract = MathExpr("subtract")
	Multiply = MathExpr("multiply")
)

func ExpressionsTest() {
	expr := mathExpression(Multiply)
	fmt.Println(expr(9, 7))

	doubleResult := double(5, 6, expr)
	fmt.Println(doubleResult)
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
		switch expr {
			case Addition:
				return add
			case Subtract:
				return subtract
			case Multiply:
				return multiply
			default:
				return func(a float64, b float64) float64 {
					return 0
				}
		}
}

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply( a float64, b float64) float64 {
	return a * b
}

func double(a float64, b float64, expr func(float64, float64) float64) float64 {
	return 2 * expr(a, b)
}
