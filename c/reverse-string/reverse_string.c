#include "reverse_string.h"

#include <stdlib.h>
#include <string.h>

char *reverse(const char *value)
{
    const size_t length = strlen(value);
    char *buffer = malloc(length + 1);

    for (size_t i = 0, j = length - 1; i < length; i++, j--) {
        buffer[i] = value[j];
    }
    buffer[length] = '\0';

    return buffer;
}
