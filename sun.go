// Package sun implements Sunrise/Sunset Algorithm
package sun

import (
	"errors"
	"math"
	"time"
)

// Error values if sun not rises or sets at speciefied date and location
var (
	ErrNoRise = errors.New("the sun never rises on this location (on the specified date)")
	ErrNoSet  = errors.New("the sun never sets on this location (on the specified date)")
)

func rad(deg float64) float64 { return deg * math.Pi / 180.0 }
func deg(rad float64) float64 { return rad * 180.0 / math.Pi }

type mode int

const (
	rising mode = iota
	setting
)

func calc(tt time.Time, lat, lon, zen float64, m mode) (time.Time, error) {
	// 1. first calculate the day of the year
	N := float64(tt.YearDay())
	// 2. convert the longitude to hour value and calculate an approximate time
	switch m {
	case rising:
		N += (6.0 - lon/15.0) / 24.0
	case setting:
		N += (18.0 - lon/15.0) / 24.0
	}
	// 3. calculate the Sun's mean anomaly
	M := 0.9856*N - 3.289
	// 4. calculate the Sun's true longitude
	L := M + 1.916*math.Sin(rad(M)) + 0.020*math.Sin(2*rad(M)) + 282.634
	// 5a. calculate the Sun's right ascension
	RA := deg(math.Atan(0.91764 * math.Tan(rad(L))))
	// 5b. right ascension value needs to be in the same quadrant as L
	RA += math.Floor(L/90.0)*90.0 - math.Floor(RA/90.0)*90.0
	// 5c. right ascension value needs to be converted into hours
	RA /= 15.0
	// 6. calculate the Sun's declination
	sinDec := 0.39782 * math.Sin(rad(L))
	cosDec := math.Cos(math.Asin(sinDec))
	// 7a. calculate the Sun's local hour angle
	cosH := (math.Cos(rad(zen)) - sinDec*math.Sin(rad(lat))) / (cosDec * math.Cos(rad(lat)))
	switch {
	case cosH > 1.0:
		return time.Time{}, ErrNoRise
	case cosH < -1.0:
		return time.Time{}, ErrNoSet
	}
	// 7b. finish calculating H and convert into hours
	H := deg(math.Acos(cosH)) / 15.0
	if m == rising {
		H = -H
	}
	// 8. calculate local mean time of rising/setting
	T := math.Mod(H+RA-0.06571*N-6.622, 24.0)
	// 9. adjust back to UTC
	UT := (T - lon/15.0) * float64(time.Hour)
	// 10. convert UT value to local time zone of latitude/longitude
	return tt.Truncate(24 * time.Hour).Add(time.Duration(UT)), nil
}

// Zenith for sunrise/sunset
type Zenith float64

// Sun's zenith for sunrise/sunset
const (
	Official     Zenith = 90.0 + 50.0/60.0
	Civil        Zenith = 96.0
	Nautical     Zenith = 102.0
	Astronomical Zenith = 108.0
)

// Rise returns a sunrise time at given time, location on given zenith
func (z Zenith) Rise(t time.Time, lat, lon float64) (time.Time, error) {
	return calc(t, lat, lon, float64(z), rising)
}

// Set returns a sunset time at given time, location on given zenith
func (z Zenith) Set(t time.Time, lat, lon float64) (time.Time, error) {
	return calc(t, lat, lon, float64(z), setting)
}

// Rise returns a sunrise time at given time, location on official zenith
func Rise(t time.Time, lat, lon float64) (time.Time, error) {
	return Official.Rise(t, lat, lon)
}

// Set returns a sunset time at given time, location on official zenith
func Set(t time.Time, lat, lon float64) (time.Time, error) {
	return Official.Set(t, lat, lon)
}
