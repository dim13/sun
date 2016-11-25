package sun

import (
	"testing"
	"time"
)

const testDate = "2009-11-10T23:00:00Z"

var testCases = []struct {
	place           string
	lat, lon        float64
	date, rise, set string
	err             error
}{
	{"Berlin", 52.5200, 13.4050, testDate,
		"2009-11-10T07:18:22+01:00", "2009-11-10T16:21:24+01:00", nil},
	{"NewYork", 40.7300, -73.9352, testDate,
		"2009-11-10T06:36:25-05:00", "2009-11-10T16:42:26-05:00", nil},
	{"Sydney", -33.8679, 151.2073, testDate,
		"2009-11-10T05:47:34+11:00", "2009-11-10T19:31:00+11:00", nil},
	{"Honolulu", 21.3069, -157.8583, testDate,
		"2009-11-10T06:39:19-10:00", "2009-11-10T17:51:16-10:00", nil},
	{"Johannesburg", -26.2041, 28.0473, testDate,
		"2009-11-10T05:13:05+02:00", "2009-11-10T18:30:42+02:00", nil},
	// Odd places
	{"null", 0.0, 0.0, testDate,
		"2009-11-10T05:40:25Z", "2009-11-10T17:47:27Z", nil},
	{"Kiritimati", 1.8721, -157.4278, testDate,
		"2009-11-11T06:12:29+14:00", "2009-11-11T18:14:52+14:00", nil},
	{"Baker Island", 0.1936, -176.4769, testDate,
		"2009-11-10T05:26:37-12:00", "2009-11-10T17:33:10-12:00", nil},
	// Errors
	{"North Pol", 90.0, 0.0, testDate,
		"0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z", ErrNoRise},
	{"South Pol", -90.0, 0.0, testDate,
		"0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z", ErrNoSet},
}

func TestRise(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.place, func(t *testing.T) {
			date, err := time.Parse(time.RFC3339, tc.date)
			if err != nil {
				t.Error(err)
			}
			r, err := Rise(date, tc.lat, tc.lon)
			if err != tc.err {
				t.Errorf("got %v; want %v", err, tc.err)
			}
			sr, err := time.Parse(time.RFC3339, tc.rise)
			if err != nil {
				t.Fatal(err)
			}
			if r = r.Truncate(time.Second); !r.Equal(sr) {
				t.Errorf("got %v; want %v", r, sr.UTC())
			}
		})
	}
}

func TestSet(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.place, func(t *testing.T) {
			date, err := time.Parse(time.RFC3339, tc.date)
			if err != nil {
				t.Error(err)
			}
			s, err := Set(date, tc.lat, tc.lon)
			if err != tc.err {
				t.Errorf("got %v; want %v", err, tc.err)
			}
			ss, err := time.Parse(time.RFC3339, tc.set)
			if err != nil {
				t.Fatal(err)
			}
			if s = s.Truncate(time.Second); !s.Equal(ss) {
				t.Errorf("got %v; want %v", s, ss.UTC())
			}
		})
	}
}

func BenchmarkRise(b *testing.B) {
	for _, tc := range testCases {
		b.Run(tc.place, func(b *testing.B) {
			date, err := time.Parse(time.RFC3339, tc.date)
			if err != nil {
				b.Error(err)
			}
			for i := 0; i < b.N; i++ {
				Rise(date, tc.lat, tc.lon)
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	for _, tc := range testCases {
		b.Run(tc.place, func(b *testing.B) {
			date, err := time.Parse(time.RFC3339, tc.date)
			if err != nil {
				b.Error(err)
			}
			for i := 0; i < b.N; i++ {
				Set(date, tc.lat, tc.lon)
			}
		})
	}
}
