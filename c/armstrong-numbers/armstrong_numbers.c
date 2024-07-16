#include "armstrong_numbers.h"

#include <limits.h>
#include <math.h>
#include <stdio.h>
#include <string.h>

#define BUFFER_SIZE (sizeof(int) * CHAR_BIT / 3 + 2)

bool is_armstrong_number(int candidate)
{
    char buffer[BUFFER_SIZE];

    // convert the int to a c-string
    const int length = sprintf(buffer, "%d", candidate);

    int sum = 0;
    for (char *p = buffer; *p; p++) // for each digit
        sum += pow(
            *p - '0', // digit
            length    // to the power of length
        );

    return sum == candidate;
}
