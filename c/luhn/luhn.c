#include "luhn.h"

#include <ctype.h>

#define BUFFER_SIZE 100

bool luhn(const char *num)
{
    // hacky way to return false on single character strings
    if (*(num + 1) == '\0') {
        return false;
    }

    char buffer[BUFFER_SIZE];

    // copy only digits into a c-string buffer
    char *p;
    for(p = buffer; *num; num++) {
        if (isdigit(*num)) {
            *p = *num;
        }
    }
    buffer[p - buffer] = '\0';

    // in reverse, iterate across every other character
    for (p -= 2; p >= buffer; p -= 2) {
        int doubled = 2 * (*p - '0'); // double the digit

        if (doubled > 9) {
            doubled -= 9;
        }

        // replace old digit w/ modified value
        *p = (char)(doubled + '0');
    }

    // sum each of the digits in the buffer
    int sum = 0;
    for (char *q = buffer; *q; q++) {
        sum += *q - '0';
    }

    return sum % 10 == 0;
}
