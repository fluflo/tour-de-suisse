package routes

type Location struct {
	name                string
	longitude, latitude float64
	description         string
	accessbileByCar     bool
	accessbileByTrain   bool
	closestLocation     *Location
	averageParkingFees  float64 // average parking fees for this location
}

/**
 * Creates a new Location
 * @param name the name of the location
 * @param longitude the longitude of the location
 * @param latitude the latitude of the location
 * @param description the description of the location
 * @param accessbileByCar Is the location accessible by cars?
 * @param accessbileByTrain Is the location accessible by trains?
 * @param closestLocation the closest location accessible by car to this location
 * @param averageParkingFees the average parking fees for this location per day
 */
func NewLocation(
	name string,
	longitude, latitude float64,
	description string,
	accessibleByCar bool,
	accessibleByTrain bool,
	closestLocation *Location,
	averageParkingFees float64,
) *Location {
	return &Location{
		name:               name,
		longitude:          longitude,
		latitude:           latitude,
		description:        description,
		accessbileByCar:    accessibleByCar,
		accessbileByTrain:  accessibleByTrain,
		closestLocation:    closestLocation,
		averageParkingFees: averageParkingFees,
	}
}
