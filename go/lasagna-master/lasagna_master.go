// The lasagna package provides utilities for 
package lasagna

// The PreparationTime function calculates the time to prepare a lasagna based
// on the layers and preparation time per layer as provided by the caller.
func PreparationTime(layers []string, prepTimePerLayer int) int {
    if prepTimePerLayer == 0 {
        prepTimePerLayer = 2
    }

    return prepTimePerLayer * len(layers) 
}

// The Quantities function calculates the quantity of noodles and sauce needed
// based on the layers provided by the caller.
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0

    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }

    return noodles, sauce
}

// The AddSecretIngredient function replaces the last value of the first slice
// with the last value of the second slice.
func AddSecretIngredient(friend []string, mine []string) {
    mine[len(mine) - 1] = friend[len(friend) - 1]
}

// The ScaleRecipe function scales each of the quantities values by the scale
// value provided by the caller.
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := []float64{}

    scale := float64(portions) / 2
    for _, quantity := range quantities {
        scaledQuantities = append(scaledQuantities, quantity * scale)
    }

    return scaledQuantities
}
