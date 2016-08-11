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

func LatLon(lat, lon float64) string {
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

	now := time.Now().Add(time.Duration(*day) * time.Hour * 24)

	zone := latlong.LookupZoneName(*lat, *lon)
	if zone != "" {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			log.Fatal(err)
		}
		now = now.In(loc)
	}

	fmt.Println("location", LatLon(*lat, *lon))
	fmt.Println("timezone", zone)

	r, err := sun.Rise(now, *lat, *lon, sun.Official)
	if err != nil {
		log.Fatal(err)
	}

	s, err := sun.Set(now, *lat, *lon, sun.Official)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sunrise ", r.Format(time.Stamp))
	fmt.Println("sunset  ", s.Format(time.Stamp))
}
