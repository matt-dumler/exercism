#include "eliuds_eggs.h"

/*
 * egg_count: returns the number of 1's in the binary representation of an
 *     integer.
 */
unsigned int egg_count(unsigned int input)
{
    unsigned int count = 0;
    while (input > 0) {
        // Increment the counter if the least significant bit is one.
        if (input & 1) {
            count++;
        }
        // Shift all bits to the right one, so we can process the next bit.
        input >>= 1;
    }

    return count;
}
