package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the expression is not valid
	errorInvalidInput = errors.New("input expression is not valid(expecting characters, that are not numbers, +, - or whitespace)")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	if input == "" {
		return "", fmt.Errorf("err: %v", errorEmptyInput)
	}

	// // check if valid
	// for _, v := range input {
	// 	if !strings.ContainsRune("0123456789+-", v) {
	// 		err = errorInvalidInput
	// 		return "", err
	// 	}
	// }

	// getting rid of whitespace
	input = strings.ReplaceAll(input, " ", "")

	delimit := strings.LastIndexAny(input, "+-")
	if delimit == -1 {
		return "", fmt.Errorf("err: %v", errorNotTwoOperands)
	}
	op1Str := input[0:delimit]
	op2Str := input[delimit:]

	op1, err := strconv.ParseInt(op1Str, 10, 64)
	if err != nil {
		return "", fmt.Errorf("err: %v", errorInvalidInput)
	}
	op2, err := strconv.ParseInt(op2Str, 10, 64)
	if err != nil {
		return "", fmt.Errorf("err: %v", errorInvalidInput)
	}

	output = strconv.FormatInt(op1+op2, 10)

	return output, nil
}
