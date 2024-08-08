package isogram

import "unicode"

func IsIsogram(word string) bool {
    letters := map[rune]bool{}

    for _, r := range word {

        r = unicode.ToLower(r)
        if unicode.IsLetter(r) {

            _, occured := letters[r]
            if occured {
                return false
            } else {
                letters[r] = true
            }
        }
    }

    return true
}
