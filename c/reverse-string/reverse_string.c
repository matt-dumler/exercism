#include "reverse_string.h"

#include <stdlib.h>
#include <string.h>

char *reverse(const char *value)
{
    const size_t length = strlen(value);

    char *buffer = malloc(length + 1);
    buffer[length] = '\0';

    for (char *p = buffer + (length - 1); *value; value++)
        *(p--) = *value;

    return buffer;
}
