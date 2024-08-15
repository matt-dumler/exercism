package phonenumber

import (
    "fmt"
    "regexp"
    "unicode"
)

func onlyDigits(input string) string {
    digits := make([]rune, 0)

    for _, r := range input {
        if unicode.IsDigit(r) {
            digits = append(digits, r)
        }
    }

    return string(digits)
}

func Number(phoneNumber string) (string, error) {
    number := onlyDigits(phoneNumber)

    re := regexp.MustCompile(`^(?:1[2-9]\d{2}[2-9]\d{6}|[2-9]\d{2}[2-9]\d{6})$`)
    isNANP := re.MatchString(number)

    if isNANP {
        if len(number) == 11 {
            number = number[1:]
        }

        return number, nil
    } else {
        return "", fmt.Errorf("%s is not a valid NANP", phoneNumber)
    }
}

func AreaCode(phoneNumber string) (string, error) {
    number, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
    n, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}
