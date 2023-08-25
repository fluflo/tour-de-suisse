package cars

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/fluflo/tour-de-suisse/routes"
	geojson "github.com/paulmach/go.geojson"
)

func GetCarsRoute() {
	trip := routes.CreateRoutes()[0]
	apiKey := os.Getenv("TOURDESUISSE_GEOAPI_KEY")
	url := fmt.Sprintf("https://api.geoapify.com/v1/routing?waypoints=%s&mode=drive&apiKey=%s", trip.GetLongituteAndLatitudesOfPathForCarRoute(), apiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fc, err := geojson.UnmarshalFeatureCollection(body)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	distance, err := getFloat(fc.Features[0].Properties["distance"])

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	car := NewCar("Golf", 5.8)
	gasoline := NewGasoline("95", 1.84)

	totalConsumption := car.CalculateTotalConsumption(distance / 1000)

	fmt.Printf("Trip distance: %f km\n", distance/1000)
	fmt.Printf("Trip locations: %s\n", trip.GetLocationNames())

	fmt.Printf("Total consumption: %f liters\nPrice for distance: %f CHF\n", totalConsumption, totalConsumption*gasoline.price)
	fmt.Printf("Total parking fees: %f CHF\n", trip.GetTotalParkingFees())
	fmt.Printf("Total fees: %f CHF\n", trip.GetTotalParkingFees()+(totalConsumption*gasoline.price))

}

func getFloat(unk interface{}) (float64, error) {
	var floatType = reflect.TypeOf(float64(0))
	var stringType = reflect.TypeOf("")

	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("can't convert %v to float64", v.Type())
		}
	}
}
