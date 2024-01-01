package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery: 100,
		batteryDrain: batteryDrain,
		speed: speed,
		distance: 0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	new_car:=  NewCar(car.speed, car.batteryDrain)
	if car.battery >= car.batteryDrain{
		new_car.battery = car.battery - car.batteryDrain
		new_car.distance = car.speed + car.distance
	} else {
		new_car.battery = car.battery
		new_car.distance = car.distance
	}
	
	return new_car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	speed:= car.speed
	dist:=track.distance
	freq:= dist/speed
	required_battery:= freq*car.batteryDrain
	if car.battery >= required_battery {
		return true
	}
	return false
}
