#include "space_age.h"

#define ERROR_VALUE -1.0

const float MERCURY_YEAR = 0.2408467,
      VENUS_YEAR = 0.61519726,
      MARS_YEAR = 1.8808158,
      JUPITER_YEAR = 11.862615,
      SATURN_YEAR = 29.447498,
      URANUS_YEAR = 84.016846,
      NEPTUNE_YEAR = 164.79132;

const int SECONDS_A_YEAR = 31557600;

float age(planet_t planet, int64_t seconds)
{
    const int YEARS_OLD = seconds / SECONDS_A_YEAR;
    switch(planet) {
        case MERCURY:
            return YEARS_OLD / MERCURY_YEAR;
            break;
        case VENUS:
            return YEARS_OLD / VENUS_YEAR;
            break;
        case EARTH:
            return YEARS_OLD;
            break;
        case MARS:
            return YEARS_OLD / MARS_YEAR;
            break;
        case JUPITER:
            return YEARS_OLD / JUPITER_YEAR;
            break;
        case SATURN:
            return YEARS_OLD / SATURN_YEAR;
            break;
        case URANUS:
            return YEARS_OLD / URANUS_YEAR;
            break;
        case NEPTUNE:
            return YEARS_OLD / NEPTUNE_YEAR;
            break;
        default:
            return ERROR_VALUE;
    }
}
