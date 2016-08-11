package sun

import (
	"testing"
	"time"
)

const (
	lat = 52.5200
	lon = 13.4050
)

func TestDayOfYear(t *testing.T) {
	d := dayOfYear(time.Now())
	t.Log(d)
}

func TestRise(t *testing.T) {
	r, err := Rise(time.Now(), lat, lon)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestSet(t *testing.T) {
	s, err := Set(time.Now(), lat, lon)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
