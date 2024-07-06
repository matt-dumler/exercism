#include "resistor_color_duo.h"

#define TENS_PLACE 10

// color_code: converts two resistor colors to a numerical value
uint8_t color_code(resistor_band_t colors[])
{
    return (colors[0] * TENS_PLACE) + colors[1];
}
