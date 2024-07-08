#include "binary.h"

// convert: converts a binary value as a string to an integer.
int convert(const char *input)
{
    /*
     * Capture the length of the binary string (not the C string).
     *
     * Calculate the value of the left most bit's place. This will be a multiple
     *     of two.
     */
    const int length = strlen(input);
    int place = pow(2, length - 1), sum = 0;

    /* As we loop through the binary string we will reduce the place value by
     *     half each step, and if the place in the binary string has a '1'
     *     character we will add the current place value to a running sum. The
     *     total sum will be the integer equivalent of the binary string.
     *
     *  Return the error value if any invalid characters are encountered.
     */
    for (int i = 0; i < length; i++) {
        const char c = input[i];

        if (c == '1') {
            sum += place;
        } else if (c != '0') {
            return INVALID;
        }

        place /= 2;
    }

    return sum;
}
