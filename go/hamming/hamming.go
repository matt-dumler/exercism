package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("a and b must be of equal length")
    }

    dist := 0
    runes := []rune(b)
    for i, c := range a {
        if runes[i] != c {
            dist++
        }
    }

    return dist, nil
}
