package sun

import (
	"testing"
	"time"
)

const (
	lat = 52.5200
	lon = 13.4050
)

var (
	when, _ = time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	rise, _ = time.Parse(time.RFC3339, "2009-11-10T06:18:22Z")
	set, _  = time.Parse(time.RFC3339, "2009-11-10T15:21:24Z")
)

func TestDayOfYear(t *testing.T) {
	d := dayOfYear(when)
	if d != 314 {
		t.Errorf("expected 314, got %v", d)
	}
}

func TestRise(t *testing.T) {
	r, err := Rise(when, lat, lon)
	if err != nil {
		t.Error(err)
	}
	if r.Truncate(time.Second) != rise {
		t.Errorf("expected %v, got %v", rise, r)
	}
}

func TestSet(t *testing.T) {
	s, err := Set(when, lat, lon)
	if err != nil {
		t.Error(err)
	}
	if s.Truncate(time.Second) != set {
		t.Errorf("expected %v, got %v", set, s)
	}
}
