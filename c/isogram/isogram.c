#include "isogram.h"

#include <ctype.h>
#include <stddef.h>

bool is_isogram(const char phrase[])
{
    if (phrase == NULL) {
        return false;
    }

    /*
     * Mask the bits corresponding to each letter's position in the alphabet.
     *
     * TODO: Update this comment to reflect what is actually being done.
     */
    int bitfield = 0;
    for (char c; (c = tolower(*phrase)); phrase++) {
        if (isalpha(c)) {
            const int mask = 1 << (c - 'a');
            if (bitfield & mask) {
                return false;
            } else {
                bitfield |= mask;
            }
        }
    }

    return true;
}
