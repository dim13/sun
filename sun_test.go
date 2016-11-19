package sun

import (
	"testing"
	"time"
)

const when = "2009-11-10T23:00:00Z"

var testCases = []struct {
	place     string
	lat, lon  float64
	rise, set string
}{
	{"Berlin", 52.5200, 13.4050,
		"2009-11-10T07:18:22+01:00", "2009-11-10T16:21:24+01:00"},
	{"NewYork", 40.7300, -73.9352,
		"2009-11-10T06:36:25-05:00", "2009-11-10T16:42:26-05:00"},
	{"Sydney", -33.8679, 151.2073,
		"2009-11-10T05:47:34+11:00", "2009-11-10T19:31:00+11:00"},
	{"Honolulu", 21.3069, -157.8583,
		"2009-11-10T06:39:19-10:00", "2009-11-10T17:51:16-10:00"},
}

func testDate(t *testing.T) time.Time {
	date, err := time.Parse(time.RFC3339, when)
	if err != nil {
		t.Fatal(err)
	}
	return date
}

func TestRise(t *testing.T) {
	date := testDate(t)
	for _, tc := range testCases {
		r, err := Rise(date, tc.lat, tc.lon)
		if err != nil {
			t.Error(err)
		}
		sr, err := time.Parse(time.RFC3339, tc.rise)
		if err != nil {
			t.Fatal(err)
		}
		if r = r.Truncate(time.Second); !r.Equal(sr) {
			t.Errorf("%v: expected %v, got %v",
				tc.place, sr.UTC(), r)
		}
	}
}

func TestSet(t *testing.T) {
	date := testDate(t)
	for _, tc := range testCases {
		s, err := Set(date, tc.lat, tc.lon)
		if err != nil {
			t.Error(err)
		}
		ss, err := time.Parse(time.RFC3339, tc.set)
		if err != nil {
			t.Fatal(err)
		}
		if s = s.Truncate(time.Second); !s.Equal(ss) {
			t.Errorf("%v: expected %v, got %v",
				tc.place, ss.UTC(), s)
		}
	}
}

func TestNoRise(t *testing.T) {
	date := testDate(t)
	s, err := Set(date, 90.0, 0.0)
	if err != ErrNoRise {
		t.Errorf("expected %v, got %v", ErrNoRise, s)
	}
}

func TestNoSet(t *testing.T) {
	date := testDate(t)
	s, err := Set(date, -90.0, 0.0)
	if err != ErrNoSet {
		t.Errorf("expected %v, got %v", ErrNoSet, s)
	}
}
