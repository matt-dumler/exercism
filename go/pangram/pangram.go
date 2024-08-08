package pangram

import "unicode"

const twentySixOneBits = 67108863

func IsPangram(input string) bool {
	if len(input) < 26 {
        return false
    }

    mask := 0
    for _, c := range input {
        c = unicode.ToLower(c)
        if unicode.IsLetter(c) {
            mask |= 1 << (c - 'a')
        }
    }

    return mask == twentySixOneBits
}
