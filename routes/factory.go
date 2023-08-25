package routes

var zurich = NewLocation("Zurich", 8.564572, 47.451542, "Zurich is not the capital city of Switzerland.", true, true, nil, 40.0)
var bern = NewLocation("Bern", 7.451123, 46.947968, "Bern is the capital city of Switzerland.", true, true, nil, 30.0)
var interlaken = NewLocation("Interlaken", 7.850000, 46.683334, "interlaken is the capital city of Switzerland.", true, true, nil, 10.0)
var lauterbrunnen = NewLocation("Lauterbrunnen", 7.900000, 46.599998, "Lauterbrunnen is a nice town with lots of waterfalls.", true, true, nil, 18.0)
var wengen = NewLocation("Wengen", 7.915562, 46.609437, "Wengen is a nice town without cars", false, true, lauterbrunnen, 0.0)
var lucerne = NewLocation("Lucernce", 8.309307, 47.050168, "Lucernce is the capital city of Switzerland.", true, true, nil, 30.0)

func CreateRoutes() []Route {
	trips := make([]Route, 0)
	trips = append(trips, *NewRoute("Short trip", []Location{*zurich, *bern, *interlaken, *lauterbrunnen, *wengen, *interlaken, *lucerne, *zurich}, "A short trip to Berne, Interlaken and Lucerne", 5))
	return trips
}
