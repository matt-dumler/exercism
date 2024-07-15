#include "gigasecond.h"

void gigasecond(time_t input, char *output, size_t size)
{
    const time_t after = input + 1000000000;
    const struct tm *utc = gmtime(&after);
    strftime(output, size, "%Y-%m-%d %H:%M:%S", utc);
}
