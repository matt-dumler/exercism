package wordcount

import (
    "strings"
    "unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    freq := Frequency{}

    withoutCommas := strings.Join(strings.Split(phrase, ","), " ")
    for _, word := range strings.Fields(withoutCommas) {
        word = strings.ToLower(word)
        word = strings.TrimFunc(word, func(r rune) bool {
            return !unicode.IsLetter(r) && !unicode.IsDigit(r)
        })

        if len(word) > 0 {
            freq[word]++
        }
    }

    return freq
}
