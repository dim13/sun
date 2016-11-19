package sun

import (
	"testing"
	"time"
)

const when = "2009-11-10T23:00:00Z"

var testCases = []struct {
	lat, lon  float64
	rise, set string
}{
	{+52.5200, +13.4050, "2009-11-10T07:18:22+01:00", "2009-11-10T16:21:24+01:00"},
	{+40.7300, -73.9352, "2009-11-10T06:36:25-05:00", "2009-11-10T16:42:26-05:00"},
	{-33.8679, 151.2073, "2009-11-10T05:47:34+11:00", "2009-11-10T19:31:00+11:00"},
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
		if !r.Truncate(time.Second).Equal(sr) {
			t.Errorf("expected %v, got %v", sr, r)
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
		if !s.Truncate(time.Second).Equal(ss) {
			t.Errorf("expected %v, got %v", ss, s)
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
