#include "leap.h"

// leap_year: given a year return a boolean indicating if it is a leap year
bool leap_year(int year)
{
	/*
	 * A year is leap year if it is evenly divisible by 4 and not 100. If a
	 * 	year is evenly divisible by both 4 and 100, it is a leap year if
	 * 	is is also evenly divisible by 400.
	 */
	return (year%4 == 0 && year%100 != 0) || year%400 == 0;
}
