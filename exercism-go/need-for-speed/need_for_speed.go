package speed

// Car: Cars start with full (100%) batteries. Each time you drive the car
// using the remote control, it covers the car's speed in meters and decreases
// the remaining battery percentage by its battery drain.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// Track: Each race track has its own distance. Cars are tested by checking if
// they can finish the track without running out of battery.
type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car.distance >= track.distance
}
