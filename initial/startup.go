package initial

import (
	"fmt"
	"os"
)

func PreStartupChecks() error {
	geoApiKey := os.Getenv("TOURDESUISSE_GEOAPI_KEY")

	if len(geoApiKey) == 0 {
		return fmt.Errorf("can't find environment value TOURDESUISSE_GEOAPI_KEY")
	}
	return nil
}
