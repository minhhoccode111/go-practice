package jedlik

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() (result string) {
	result = fmt.Sprintf("Driven %v meters", car.distance)
	return
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() (result string) {
	result = fmt.Sprintf("Battery at %v%%", car.battery)
	return
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	for car.battery > 0 {
		car.Drive()
	}
	return car.distance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
