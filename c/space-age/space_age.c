#include "space_age.h"

#define ERROR_VALUE -1.0

static const float PLANET_YEAR_RATIOS[] = {
    0.2408467,  // Mercury year
    0.61519726, // Venus year
    1.0,        // Earth year
    1.8808158,  // Mars year
    11.862615,  // Jupiter year
    29.447498,  // Saturn year
    84.016846,  // Uranus year
    164.79132   // Neptune year
};

static const int SECONDS_A_YEAR = 31557600;

float age(planet_t planet, int64_t seconds)
{
    if (planet >= 0 && planet < 8) {
        return seconds / SECONDS_A_YEAR / PLANET_YEAR_RATIOS[planet];
    } else {
        return ERROR_VALUE;
    }
}
