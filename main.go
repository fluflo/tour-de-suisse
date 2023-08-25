package main

import (
	"github.com/fluflo/tour-de-suisse/cars"
	"github.com/fluflo/tour-de-suisse/initial"
)

func main() {
	err := initial.PreStartupChecks()
	if err != nil {
		panic(err)
	}
	cars.GetCarsRoute()
}
