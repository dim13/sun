// Show sunrise and sunset on given location
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/latlong"

	"dim13.org/sun"
)

var (
	lat = flag.Float64("lat", 52.5200, "latitude")
	lon = flag.Float64("lon", 13.4050, "longitude")
	day = flag.Int("days", 0, "offset")
)

func latLon(lat, lon float64) string {
	LA := 'N'
	LO := 'E'
	if lat < 0.0 {
		lat = -lat
		LA = 'S'
	}
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

	fmt.Println("location", latLon(*lat, *lon))
	if zone == "" {
		zone = "UTC"
	}
	fmt.Println("timezone", zone)

	r, err := sun.Rise(now, *lat, *lon)
	if err != nil {
		log.Fatal(err)
	}

	s, err := sun.Set(now, *lat, *lon)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sunrise ", r.Format(time.Stamp))
	fmt.Println("sunset  ", s.Format(time.Stamp))
}
