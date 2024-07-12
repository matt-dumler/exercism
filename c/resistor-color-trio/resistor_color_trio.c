#include "resistor_color_trio.h"

#define TENS_PLACE 10
#define HUNDREDS_PLACE 100

// color_code: converts three resistor colors to a numerical value
uint8_t color_code(resistor_band_t colors[])
{
    return (colors[0] * HUNDREDS_PLACE) + (colors[1] * TENS_PLACE) + colors[2];
}
