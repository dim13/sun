// Show sunrise and sunset on given location
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/latlong"

	"github.com/dim13/sun"
)

var (
	lat = flag.Float64("lat", 52.5200, "latitude")
	lon = flag.Float64("lon", 13.4050, "longitude")
	day = flag.Int("days", 0, "offset")
)

func latLon(lat, lon float64) string {
	LA := 'N'
	if lat < 0.0 {
		lat = -lat
		LA = 'S'
	}
	LO := 'E'
	if lon < 0.0 {
		lon = -lon
		LO = 'W'
	}
	return fmt.Sprintf("%v°%c %v°%c", lat, LA, lon, LO)
}

func main() {
	flag.Parse()

	zone := latlong.LookupZoneName(*lat, *lon)
	loc, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().In(loc).Add(time.Duration(*day) * time.Hour * 24)

	sunRise, err := sun.Rise(now, *lat, *lon)
	if err != nil {
		log.Fatal(err)
	}

	sunSet, err := sun.Set(now, *lat, *lon)
	if err != nil {
		log.Fatal(err)
	}

	dayDuration := sunSet.Sub(sunRise).Truncate(time.Second)

	fmt.Println("location    ", latLon(*lat, *lon))
	fmt.Println("timezone    ", loc)
	fmt.Println("sun rise    ", sunRise.Format(time.Stamp))
	fmt.Println("sun set     ", sunSet.Format(time.Stamp))
	fmt.Println("day duration", dayDuration)
}
