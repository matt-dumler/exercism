#include "hamming.h"

#define ERROR_VALUE -1

/*
 * compute: returns the Hamming distance of two DNA strands provided as
 *     strings.
 *
 * For more information on Hamming distance, see the Wikipedia page:
 *     https://en.wikipedia.org/wiki/Hamming_distance
 */
int compute(const char *lhs, const char *rhs)
{
    /*
     * Capture the length of each string, and ensure they are equal lengths.
     */
    const int len_lhs = strlen(lhs),
          len_rhs = strlen(rhs);

    if (len_lhs != len_rhs) {
        return ERROR_VALUE;
    }

    /*
     * Compare each string's characters, and total the differences.
     */
    int distance = 0;
    for (int i = 0; i < len_lhs; i++) {
        if (lhs[i] != rhs[i]) {
            distance++;
        }
    }

    return distance;
}
