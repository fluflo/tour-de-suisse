package cars

type Car struct {
	name        string
	consumption float64 // l/100km
}

func NewCar(name string, consumption float64) *Car {
	return &Car{name, consumption}
}

/**
 * Calculate the total gasoline consumption of a car for a given distance.
 * @param distance distance in km
 * @return the total gasoline consumption in l
 */
func (c *Car) CalculateTotalConsumption(distance float64) float64 {
	return distance * c.consumption / 100
}
