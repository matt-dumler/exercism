package summultiples

func SumMultiples(limit int, divisors ...int) int {
    multiples := map[int]bool{}

    for _, divisor := range divisors {
        if divisor != 0 {
            for i := 1; i * divisor < limit; i++ {
                multiples[i * divisor] = true
            }
        }
    }

    sum := 0
    for multiple := range multiples {
        sum += multiple
    }

    return sum
}
