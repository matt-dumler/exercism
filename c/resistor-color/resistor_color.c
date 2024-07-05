#include "resistor_color.h"

resistor_band_t color_arr[] = { COLORS };

resistor_band_t* colors(void)
{
	return color_arr;
}

uint8_t color_code(resistor_band_t color)
{
	return color;
}
