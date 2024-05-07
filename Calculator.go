//extensionfthecircle@gmail.com

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(romanA, romanB string) (int, int, error) {
	romanToArabicMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	arabicA, okA := romanToArabicMap[romanA]
	arabicB, okB := romanToArabicMap[romanB]
	if !okA || !okB {
		return 0, 0, errors.New("invalid numeral. ONLY Roman or ONLY Arabic numberals form 1 to 10")
	}
	return arabicA, arabicB, nil
}

func arabicToRoman(arabic int) string {
	if arabic < 1 {
		return "Roman numeral result < 1"
	}
	//range для map не имеет порядка перебора, поэтому slice:(
	arabicToRomanSlice := []struct {
		arabic int
		roman  string
	}{
		{100, "C"}, {90, "XC"}, {80, "LXXX"}, {70, "LXX"}, {60, "LX"},
		{50, "L"}, {40, "XL"}, {30, "XXX"}, {20, "XX"}, {10, "X"},
		{9, "IX"}, {8, "VIII"}, {7, "VII"}, {6, "VI"}, {5, "V"},
		{4, "IV"}, {3, "III"}, {2, "II"}, {1, "I"},
	}

	roman := ""
	for _, current := range arabicToRomanSlice {
		for arabic >= current.arabic && arabic > 0 {
			roman += current.roman
			arabic -= current.arabic
		}
	}
	return roman
}
func inputConverter(expression string) (int, string, int, bool, error) {
	operators := []string{"+", "-", "*", "/"}
	romanType := false
	for _, operator := range operators {
		if strings.Contains(expression, operator) {
			splittedExpression := strings.Split(expression, operator)
			if len(splittedExpression) == 2 {
				var A, B int
				var errA, errB error

				A, errA = strconv.Atoi(splittedExpression[0])
				B, errB = strconv.Atoi(splittedExpression[1])
				if errA != nil || errB != nil {
					var err error
					A, B, err = romanToArabic(splittedExpression[0], splittedExpression[1])
					if err != nil {
						return 0, "", 0, false, err
					}
					romanType = true
				}
				if (A > 10) || (A <= 0) || (B > 10) || (B <= 0) {
					return 0, "", 0, false, errors.New("Only numerals from 1 to 10")
				}
				return A, operator, B, romanType, nil
			} else {
				return 0, "", 0, false, errors.New("Too hard expression for me:)")
			}
		}
	}
	return 0, "", 0, false, errors.New("invalid expression format")
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
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Kata Test Task(Go Calculator)")
		fmt.Println("Enter expression:")
		expression, _ := reader.ReadString('\n')
		expression = strings.TrimSpace(expression)
		expression = strings.ReplaceAll(expression, " ", "")
		A, operator, B, romanType, err := inputConverter(expression)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		result := Calculator(A, operator, B)
		if romanType {
			fmt.Printf("Result: %s = %s\n", expression, arabicToRoman(result))
		} else {
			fmt.Printf("Result: %s = %d\n", expression, result)
		}
	}

}
