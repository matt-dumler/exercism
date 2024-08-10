package bottlesong

import (
    "fmt"
    "strings"
)

var words = [...]string{
    "no",
    "One",
    "Two",
    "Three",
    "Four",
    "Five",
    "Six",
    "Seven",
    "Eight",
    "Nine",
    "Ten",
}

func bottleWord(quantity int) string {
    if quantity == 1 {
        return "bottle"
    } else {
        return "bottles"
    }
}

func verse(bottles int) []string {
    lines := []string{}

    template := "%s green %s hanging on the wall,"
    onTheWall := fmt.Sprintf(template, words[bottles], bottleWord(bottles))
    lines = append(lines, onTheWall)
    lines = append(lines, onTheWall)

    lines = append(lines, "And if one green bottle should accidentally fall,")
    bottles--

    template = "There'll be %s green %s hanging on the wall."
    remainingBottles := fmt.Sprintf(template, strings.ToLower(words[bottles]), bottleWord(bottles))
    lines = append(lines, remainingBottles)

    return lines
}

func Recite(startBottles, takeDown int) []string {
    lines := []string{}
	for startBottles > 0 {
        lines = append(lines, verse(startBottles)...)
        startBottles--

        takeDown--
        if takeDown == 0 {
            break
        }

        lines = append(lines, "")
    }

    return lines
}
