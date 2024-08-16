package thefarm

import (
    "errors"
    "fmt"
)

var ErrNoCows = errors.New("invalid number of cows")

func DivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
    amount, err := calc.FodderAmount(numberOfCows)
    if err != nil {
        return 0, err
    }

    factor, err := calc.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return amount * factor / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
    if n < 1 {
        return 0, ErrNoCows
    }

    return DivideFood(fc, n)
}

type InvalidCowsError struct {
    num int
    msg string
}

func (err *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", err.num, err.msg)
}

func ValidateNumberOfCows(n int) error {
    if n < 0 {
        return &InvalidCowsError{n, "there are no negative cows"}
    } else if n < 1 {
        return &InvalidCowsError{n, "no cows don't need food"}
    }

    return nil
}
