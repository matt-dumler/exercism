#ifndef RESISTOR_COLOR_H
#define RESISTOR_COLOR_H

#include <stdint.h>

// single point of truth for the list of colors
#define COLORS \
	BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE

typedef enum { COLORS } resistor_band_t;

resistor_band_t* colors(void);

uint8_t color_code(resistor_band_t color);

#endif
