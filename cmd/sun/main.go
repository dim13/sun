package main

import (
	"flag"
	"fmt"
	"time"

	"dim13.org/sun"
)

var (
	lat = flag.Float64("lat", 52.5200, "latitude")
	lon = flag.Float64("lon", 13.4050, "longitude")
)

func main() {
	flag.Parse()
	now := time.Now()

	if r, err := sun.Rise(now, *lat, *lon, sun.Official); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sunrise", r)
	}

	if s, err := sun.Set(now, *lat, *lon, sun.Official); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sunset ", s)
	}
}
