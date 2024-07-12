#include "resistor_color_trio.h"

#define TENS_PLACE 10

// color_code: converts three resistor colors to a numerical value
resistor_value_t color_code(resistor_band_t colors[])
{
    struct resistor_value_t result;

    result.value = (10 * colors[0]) + colors[1];

    resistor_band_t color = colors[2];
    switch (color) {
        case BLACK:
            result.unit = OHMS;
            break;
        case BROWN:
            result.unit = OHMS;
            result.value *= 10;
            break;
        case RED:
            result.unit = KILOOHMS;
            result.value /= 10;
            break;
        case ORANGE:
            result.unit = KILOOHMS;
            break;
        case YELLOW:
            result.unit = KILOOHMS;
            result.value *= 10;
            break;
        case GREEN:
            result.unit = KILOOHMS;
            result.value *= 100;
            break;
        case BLUE:
            result.unit = MEGAOHMS;
            break;
        case VIOLET:
            result.unit = MEGAOHMS;
            result.value *= 10;
            break;
        case GREY:
            result.unit = MEGAOHMS;
            result.value *= 100;
            break;
        case WHITE:
            result.unit = GIGAOHMS;
            break;
    }

    return result;
}
