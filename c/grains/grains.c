#include "grains.h"

/*
 * See the following wikipedia article to understand how the following two
 * 	formulas are derived:
 *
 * 	https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem#Solutions
 */

uint64_t square(uint8_t index)
{
	return pow(2, index-1);
}

uint64_t total(void)
{
	return powl(2, 64) - 1;
}
