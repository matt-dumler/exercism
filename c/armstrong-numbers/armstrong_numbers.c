#include "armstrong_numbers.h"

#include <math.h>
#include <stdio.h>
#include <string.h>

#define BUFFER_SIZE 100

bool is_armstrong_number(int candidate)
{
    char buffer[BUFFER_SIZE];

    sprintf(buffer, "%d", candidate);
    const int length = strlen(buffer);

    int sum = 0;
    for (char *p = buffer; *p; p++) {
            const int digit = *p - '0';
            sum += pow(digit, length);
    }

    return sum == candidate;
}
