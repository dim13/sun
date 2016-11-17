package sun

import (
	"testing"
	"time"
)

const (
	lat  = 52.5200
	lon  = 13.4050
	when = "2009-11-10T23:00:00Z"
	rise = "2009-11-10T06:18:22Z"
	set  = "2009-11-10T15:21:24Z"
)

func TestDayOfYear(t *testing.T) {
	date, err := time.Parse(time.RFC3339, when)
	if err != nil {
		t.Fatal(err)
	}
	d := dayOfYear(date)
	if d != 314 {
		t.Errorf("expected 314, got %v", d)
	}
}

func TestRise(t *testing.T) {
	date, err := time.Parse(time.RFC3339, when)
	if err != nil {
		t.Fatal(err)
	}
	r, err := Rise(date, lat, lon)
	if err != nil {
		t.Error(err)
	}
	sr, err := time.Parse(time.RFC3339, rise)
	if err != nil {
		t.Fatal(err)
	}
	if r.Truncate(time.Second) != sr {
		t.Errorf("expected %v, got %v", rise, sr)
	}
}

func TestSet(t *testing.T) {
	date, err := time.Parse(time.RFC3339, when)
	if err != nil {
		t.Fatal(err)
	}
	s, err := Set(date, lat, lon)
	if err != nil {
		t.Error(err)
	}
	ss, err := time.Parse(time.RFC3339, set)
	if err != nil {
		t.Fatal(err)
	}
	if s.Truncate(time.Second) != ss {
		t.Errorf("expected %v, got %v", set, ss)
	}
}
