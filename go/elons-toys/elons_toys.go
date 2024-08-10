package elon

import "fmt"

func (car *Car) Drive() {
    if car.battery >= car.batteryDrain {
        car.distance += car.speed
        car.battery  -= car.batteryDrain
    }
}

func (car Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
    return car.battery >= trackDistance * car.batteryDrain / car.speed
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
