package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("n must be positive")
    }

    stepCount := 0
    for n > 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = 3 * n + 1
        }
        stepCount++
    }

    return stepCount, nil
}
