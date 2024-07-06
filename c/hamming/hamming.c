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
     * Ensure the strings are not null and equal length, and capture the value.
     */
    size_t length;
    if (!lhs || !rhs || strlen(lhs) != (length = strlen(rhs))) {
        return ERROR_VALUE;
    }

    /*
     * Compare each string's characters, and total the differences.
     */
    size_t ndifferences = 0;
    for (size_t i = 0; i < length; i++) {
        if (lhs[i] != rhs[i]) {
            ndifferences++;
        }
    }

    return ndifferences;
}
