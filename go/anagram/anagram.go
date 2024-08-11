package anagram

import "strings"

func mapChars(word string) map[rune]int {
    chars := map[rune]int{}

    for _, char := range strings.ToLower(word) {
        chars[char]++
    }

    return chars
}

func equalMaps(a, b map[rune]int) bool {
    for k, v := range a {
        if v != b[k] {
            return false
        }
    }

    return true
}

func Detect(subject string, candidates []string) []string {
    results := []string{}

    mapOfChars := mapChars(subject)
    for _, word := range candidates {
        if len(word) != len(subject) || strings.ToLower(word) == strings.ToLower(subject) {
            continue
        }

        if equalMaps(mapOfChars, mapChars(word)) {
            results = append(results, word)
        }
    }

    return results
}
