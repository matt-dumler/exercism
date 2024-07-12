// ovenTime returns the amount in minutes that the lasagna should stay in the
// oven.
consteval int ovenTime() {
    return 40;
}

/* remainingOvenTime returns the remaining
   minutes based on the actual minutes already in the oven.
*/
consteval int remainingOvenTime(const int actualMinutesInOven) {
    return ovenTime() - actualMinutesInOven;
}

/* preparationTime returns an estimate of the preparation time based on the
   number of layers and the necessary time per layer.
*/
consteval int preparationTime(const int numberOfLayers) {
    return 2 * numberOfLayers;
}

// elapsedTime calculates the total time spent to create and bake the lasagna so
// far.
consteval int elapsedTime(const int numberOfLayers, const int actualMinutesInOven) {
    return preparationTime(numberOfLayers) + actualMinutesInOven;
}
