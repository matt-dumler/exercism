package cipher

import (
    "strings"
    "unicode"
)

const alphabetSize = 26

type shift struct {
    distance int
}

type vigenere struct {
    key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
    if limit := alphabetSize - 1; distance < -limit || distance == 0 || distance > limit {
        return nil
    }

	return shift{distance}
}

func shiftRuneByN(r rune, n int) rune {
    if n < 0 {
        n += alphabetSize
    }

    index := int(r - 'a')     // The index of r in the alphabet
    moved := (index + n) % 26 // The index of r once shifted

    shifted := rune('a' + moved)

    return shifted
}

func shiftStringByN(input string, distance int) string {
    buffer := make([]rune, 0)

    input = strings.ToLower(input)
    for _, r := range input {
        if unicode.IsLetter(r) {
            r = shiftRuneByN(r, distance)
            buffer = append(buffer, r)
        }
    }

    return string(buffer)
}

func (c shift) Encode(input string) string {
    return shiftStringByN(input, c.distance)
}

func (c shift) Decode(input string) string {
    return shiftStringByN(input, -c.distance)
}

func isValidKey(key string) bool {
    otherThanA := false
    for _, r := range key {
        if !unicode.IsLetter(r) || unicode.IsUpper(r) {
            return false
        }
        if r != 'a' {
            otherThanA = true
        }
    }

    return otherThanA
}

func NewVigenere(key string) Cipher {
    ok := isValidKey(key)
    if !ok {
        return nil
    }

    return vigenere{key}
}

func vigenereShift(input, key string, modifier int) string {
    buffer := make([]rune, 0)

    input = strings.ToLower(input)

    spacer := 0
    for i, r := range input {
        i = i - spacer
        if unicode.IsLetter(r) {
            if i >= len(key) {
                i = i % len(key)
            }
            distance := int(key[i]) - 'a'
            distance *= modifier

            r = shiftRuneByN(r, distance)
            buffer = append(buffer, r)
        } else {
            spacer++
        }
    }

    return string(buffer)
}

func (v vigenere) Encode(input string) string {
    return vigenereShift(input, v.key, 1)
}

func (v vigenere) Decode(input string) string {
    return vigenereShift(input, v.key, -1)
}
