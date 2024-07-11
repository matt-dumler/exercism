#include "pangram.h"

#include <ctype.h>
#include <string.h>

#define TWENTY_SIX_ONE_BITS 67108863

bool is_pangram(const char *sentence)
{
    if (sentence == NULL) {
        return false;
    }

    /*
     * Mask the bits corresponding to each letter's position in the alphabet.
     */
    int mask = 0;
    for (char c; (c = tolower(*sentence)) != '\0'; sentence++) {
        if (isalpha(c)) {
            mask |= 1 << (c - 'a');
        }
    }

    return mask == TWENTY_SIX_ONE_BITS; // Are all 26 bits masked?
}
