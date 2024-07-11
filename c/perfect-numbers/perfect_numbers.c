#include "perfect_numbers.h"

kind classify_number(int n)
{
    if (n < 1) {
        return ERROR;
    }

    // Sum all the factors of n.
    int sum = 0;
    for (int i = 1; i < n; i++) {
        if (n%i == 0) {
            sum += i;
        }
    }

    if (sum > n) {
        return ABUNDANT_NUMBER;
    } else if (sum < n) {
        return DEFICIENT_NUMBER;
    } else {
        return PERFECT_NUMBER;
    }
}
