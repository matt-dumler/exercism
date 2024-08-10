package acronym

import (
    "strings"
    "unicode"
)

func Abbreviate(s string) string {
    abbr := []rune{}

    for _, word := range strings.Split(s, " ") {
        for _, inner := range strings.Split(word, "-") {
            for _, char := range inner {
                if unicode.IsLetter(char) {
                    abbr = append(abbr, unicode.ToUpper(char))
                    break
                }
            }
        }
    }

	return string(abbr)
}
