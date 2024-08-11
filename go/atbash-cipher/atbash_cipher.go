package atbash

import "unicode"

func Atbash(s string) string {
    counter, runes := 0, []rune{}
    for _, r := range s {
        if unicode.IsSpace(r) ||
            unicode.IsPunct(r) {
            continue
        }

        if unicode.IsLetter(r) {
            r = unicode.ToLower(r)

            if r < 'n' {
                r = 'z' - (r - 'a')
            } else {
                r = 'a' + ('z' - r)
            }
        }

        if counter > 0 &&
            counter % 5 == 0 {
            runes = append(runes, ' ')
        }
        counter++

        runes = append(runes, r)
    }

    return string(runes)
}
