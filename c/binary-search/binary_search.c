#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length)
{
    if (length) {
        const int middle = length / 2;
        if (arr[middle] == value) {
            return (arr + middle);
        }

        if (middle > 0) {
            if (value < arr[middle]) {
                return binary_search(value, arr, middle);
            } else {
                return binary_search(value, (arr + middle), (length - middle));
            }
        }
    }

    return NULL;
}
