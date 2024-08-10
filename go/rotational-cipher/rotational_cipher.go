package rotationalcipher

import "unicode"

func rotate(c rune, n int) rune {
    x := int(unicode.ToLower(c)) + n
    if x > 'z' {
        x += 'a' - 'z' - 1
    }

    if unicode.IsUpper(c) {
        return unicode.ToUpper(rune(x))
    } else {
        return rune(x)
    }
}

func RotationalCipher(plain string, shiftKey int) string {
    characters := []rune{}
    for _, c := range plain {
        if unicode.IsLetter(c) {
            c = rotate(c, shiftKey)
        }

        characters = append(characters, c)
    }

    return string(characters)
}
