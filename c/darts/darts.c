#include "darts.h"

const float INNER_RING_RADIUS = 1.0;
const float MIDDLE_RING_RADIUS = 5.0;
const float OUTER_RING_RADIUS = 10.0;

// score: returns the points value for the given landing position
uint8_t score(coordinate_t landing_position)
{
    // calculate the distance of the landing position from the center
    const float radius = hypot(landing_position.x, landing_position.y);

    if (radius <= INNER_RING_RADIUS) {
        return 10;
    } else if (radius <= MIDDLE_RING_RADIUS) {
        return 5;
    } else if (radius <= OUTER_RING_RADIUS) {
        return 1;
    } else {
        return 0;
    }
}
