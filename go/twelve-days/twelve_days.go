package twelve

import (
    "fmt"
    "strings"
)

var days = [...]string{
    "first",
    "second",
    "third",
    "fourth",
    "fifth",
    "sixth",
    "seventh",
    "eighth",
    "ninth",
    "tenth",
    "eleventh",
    "twelfth",
}

var gifts = [...]string{
    " twelve Drummers Drumming,",
    " eleven Pipers Piping,",
    " ten Lords-a-Leaping,",
    " nine Ladies Dancing,",
    " eight Maids-a-Milking,",
    " seven Swans-a-Swimming,",
    " six Geese-a-Laying,",
    " five Gold Rings,",
    " four Calling Birds,",
    " three French Hens,",
    " two Turtle Doves,",
    " a Partridge in a Pear Tree.",
}

func Verse(i int) string {
    nthDay := days[i-1]

    verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", nthDay)

    daysGifts := gifts[len(gifts)-i:] 

    for j, gift := range daysGifts {
        // Include "and" if on the last gift and not the first day
        if i > 1 && j == len(daysGifts) - 1 {
            verse += " and"
        }

        verse += gift
    }

    return verse
}

func Song() string {
    song := []string{}

    for i := range days {
        song = append(song, Verse(i+1))
    }

    return strings.Join(song, "\n")
}
