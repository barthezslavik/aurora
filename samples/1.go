package main

import (
	"errors"
	"fmt"
)

// ValidateAge checks if the age is within the specified range.
func ValidateAge(age int) (int, error) {
	if age >= 0 && age <= 150 {
		return age, nil
	}
	return 0, errors.New("age is out of range")
}

func main() {
	age, err := ValidateAge(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(age)
	}
}
