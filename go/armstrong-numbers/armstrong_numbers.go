package armstrong

import (
    "math"
    "strconv"
)

func IsNumber(n int) bool {
    digits := strconv.Itoa(n)
    exponent := float64(len(digits))

    sum := 0
    for _, r := range digits {
        digit := float64(r - '0')
        sum += int(math.Pow(digit, exponent))
    }

    return sum == n
}
