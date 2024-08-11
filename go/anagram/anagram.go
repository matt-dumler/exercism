package anagram

import (
    "sort"
    "strings"
)

func sortString(s string) string {
    chars := strings.Split(s, "")
    sort.Strings(chars)
    return strings.Join(chars, "")
}

func isAnagram(a, b string) bool {
    return a != b && sortString(a) == sortString(b)
}

func Detect(subject string, candidates []string) []string {
    results := []string{}

    subject = strings.ToLower(subject)
    for _, candidate := range candidates {
        if isAnagram(subject, strings.ToLower(candidate)) {
            results = append(results, candidate)
        }
    }

    return results
}
