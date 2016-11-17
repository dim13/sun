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

func dayOfYear(t time.Time) int {
	d, m, y := t.Day(), int(t.Month()), t.Year()
	N1 := 275 * m / 9
	N2 := (m + 9) / 12
	N3 := 1 + (y-4*(y/4)+2)/3
	return N1 - (N2 * N3) + d - 30
}

func mod(v, m float64) float64 { return math.Mod(v+m, m) }
func rad(deg float64) float64  { return deg * math.Pi / 180.0 }
func deg(rad float64) float64  { return rad * 180.0 / math.Pi }

type mode int

const (
	rising mode = iota
	setting
)

func calc(tt time.Time, lat, lon, zen float64, m mode) (time.Time, error) {
	// 1. first calculate the day of the year
	N := float64(dayOfYear(tt))
	// 2. convert the longitude to hour value and calculate an approximate time
	lonHour := lon / 15.0
	var t float64
	switch m {
	case rising:
		t = N + (6.0-lonHour)/24.0
	case setting:
		t = N + (18.0-lonHour)/24.0
	}
	// 3. calculate the Sun's mean anomaly
	M := 0.9856*t - 3.289
	// 4. calculate the Sun's true longitude
	L := mod(M+1.916*math.Sin(rad(M))+0.020*math.Sin(2*rad(M))+282.634, 360.0)
	// 5a. calculate the Sun's right ascension
	RA := mod(deg(math.Atan(0.91764*math.Tan(rad(L)))), 360.0)
	// 5b. right ascension value needs to be in the same quadrant as L
	Lquad := math.Floor(L/90.0) * 90.0
	RAquad := math.Floor(RA/90.0) * 90.0
	RA += Lquad - RAquad
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
	var H float64
	switch m {
	case rising:
		H = 360.0 - deg(math.Acos(cosH))
	case setting:
		H = deg(math.Acos(cosH))
	}
	H /= 15.0
	// 8. calculate local mean time of rising/setting
	T := H + RA - 0.06571*t - 6.622
	// 9. adjust back to UTC
	UT := mod(T-lonHour, 24.0) * float64(time.Hour)
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
