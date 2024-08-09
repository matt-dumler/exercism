package darts

import "math"

const (
    innerRing = 1.0
    middleRing = 5.0
    outerRing = 10.0
)

func Score(x, y float64) int {
    radius := math.Hypot(x, y)

    if radius <= innerRing {
        return 10
    } else if radius <= middleRing {
        return 5
    } else if radius <= outerRing {
        return 1
    } else {
        return 0
    }
}
