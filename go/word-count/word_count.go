package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    freq := Frequency{}
    regex := regexp.MustCompile(`\w+('\w+)?`)

    for _, word := range regex.FindAllString(strings.ToLower(phrase), -1) {
        freq[word]++
    }

    return freq
}
