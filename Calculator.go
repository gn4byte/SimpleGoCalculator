package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	return romanMap[roman]
}
func inputConverter(expression string) (int, string, int) {
	operators := []string{"+", "-", "*", "/"}

	for _, operator := range operators {
		if strings.Contains(expression, operator) {
			splittedExpression := strings.Split(expression, operator)
			if len(splittedExpression) == 2 {
				A, _ := strconv.Atoi(splittedExpression[0])
				B, _ := strconv.Atoi(splittedExpression[1])
				fmt.Printf("Строка \"%s\" содержит оператор \"%s\"\n", expression, operator)
				return A, operator, B
			}

		}
	}
	return 0, "", 0
}
func Calculator(A int, operator string, B int) int {
	var result int
	switch operator {
	case "+":
		result = A + B
	case "-":
		result = A - B
	case "*":
		result = A * B
	case "/":
		result = A / B
	}
	fmt.Printf("Результат операции: %d %s %d = %d\n", A, operator, B, result)
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter expression: \n(In one line)")
		expression, _ := reader.ReadString('\n')
		expression = strings.TrimSpace(expression)
		Calculator(inputConverter(expression))
	}

}
