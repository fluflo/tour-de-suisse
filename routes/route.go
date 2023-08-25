package routes

import (
	"fmt"
	"strings"
)

type Route struct {
	name        string
	path        []Location
	description string
	duration    int // days
}

func NewRoute(name string, path []Location, description string, duration int) *Route {
	return &Route{
		name:        name,
		path:        path,
		description: description,
		duration:    duration,
	}
}

func (r *Route) GetLongituteAndLatitudesOfPathForCarRoute() string {
	var sb strings.Builder
	for _, location := range r.path {
		if location.accessbileByCar {
			sb.WriteString(fmt.Sprintf("%f", location.latitude))
			sb.WriteString(",")
			sb.WriteString(fmt.Sprintf("%f", location.longitude))
			sb.WriteString("|")
		} else {
			sb.WriteString(fmt.Sprintf("%f", location.closestLocation.latitude))
			sb.WriteString(",")
			sb.WriteString(fmt.Sprintf("%f", location.closestLocation.longitude))
			sb.WriteString("|")
		}

	}
	joined := sb.String()
	joined = strings.TrimSuffix(joined, "|")
	return joined
}

func (r *Route) GetLocationNames() []string {
	var names []string
	for _, location := range r.path {
		names = append(names, location.name)
	}
	return names
}

func (r *Route) AllAccessibleByCar() bool {
	for _, location := range r.path {
		if !location.accessbileByCar {
			return false
		}
	}
	return true
}

func (r *Route) GetTotalParkingFees() float64 {
	var fees float64
	for _, location := range r.path {
		fees += location.averageParkingFees
	}
	return fees / float64(len(r.path)) * float64(r.duration)
}
