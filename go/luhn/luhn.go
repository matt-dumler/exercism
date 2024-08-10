package luhn

import (
    "strconv"
    "strings"
    "unicode"
)

func Valid(id string) bool {
    // Remove the spaces from id
    id = strings.Join(strings.Split(id, " "), "")

    if len(id) < 2 {
        return false
    }

    reversed := []rune(id)
    for i, j := 0, (len(reversed) - 1); i < j; i, j = (i + 1), (j - 1) {
        reversed[i], reversed[j] = reversed[j], reversed[i]
    }

    sum := 0
    for index, rooney := range reversed {
        if !unicode.IsDigit(rooney) {
            return false
        }

        num, _ := strconv.Atoi(string(rooney))

        if index % 2 != 0 {
            num *= 2
            if num > 9 {
                num -= 9
            }
        }

        sum += num
    }

    return sum % 10 == 0
}
