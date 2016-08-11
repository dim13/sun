package main

import (
	"flag"
	"fmt"
	"os"
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

	fail := func(err error) {
		fmt.Println(err)
		os.Exit(0)
	}

	r, err := sun.Rise(now, *lat, *lon, sun.Official)
	if err != nil {
		fail(err)
	}

	s, err := sun.Set(now, *lat, *lon, sun.Official)
	if err != nil {
		fail(err)
	}

	fmt.Println("sunrise", r)
	fmt.Println("sunset ", s)
}
