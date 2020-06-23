package sun

import (
	"testing"
	"time"
)

func parseTime(v string) time.Time {
	date, err := time.Parse(time.RFC3339, v)
	if err != nil {
		panic(err)
	}
	return date
}

func parseDuration(v string) time.Duration {
	d, err := time.ParseDuration(v)
	if err != nil {
		panic(err)
	}
	return d
}

var testCases = []struct {
	place    string
	lat, lon float64
	date     time.Time
	rise     time.Time
	set      time.Time
	err      error
	day      time.Duration
}{
	{
		place: "Berlin",
		lat:   52.5200,
		lon:   13.4050,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T07:18:22+01:00"),
		set:   parseTime("2009-11-10T16:21:24+01:00"),
		day:   parseDuration("9h3m1s"),
	},
	{
		place: "London",
		lat:   51.5074,
		lon:   0.1278,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T07:08:09Z"),
		set:   parseTime("2009-11-10T16:17:52Z"),
		day:   parseDuration("9h9m43s"),
	},
	{
		place: "New York",
		lat:   40.7300,
		lon:   -73.9352,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T06:36:25-05:00"),
		set:   parseTime("2009-11-10T16:42:26-05:00"),
		day:   parseDuration("10h6m"),
	},
	{
		place: "Tokyo",
		lat:   35.6895,
		lon:   139.6917,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T06:11:18+09:00"),
		set:   parseTime("2009-11-10T16:38:31+09:00"),
		day:   parseDuration("10h27m12s"),
	},
	{
		place: "Sydney",
		lat:   -33.8679,
		lon:   151.2073,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T05:47:34+11:00"),
		set:   parseTime("2009-11-10T19:31:00+11:00"),
		day:   parseDuration("13h43m25s"),
	},
	{
		place: "Honolulu",
		lat:   21.3069,
		lon:   -157.8583,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T06:39:19-10:00"),
		set:   parseTime("2009-11-10T17:51:16-10:00"),
		day:   parseDuration("11h11m56s"),
	},
	{
		place: "Johannesburg",
		lat:   -26.2041,
		lon:   28.0473,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T05:13:05+02:00"),
		set:   parseTime("2009-11-10T18:30:42+02:00"),
		day:   parseDuration("13h17m37s"),
	},
	// Odd places
	{
		place: "Null Island",
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T05:40:25Z"),
		set:   parseTime("2009-11-10T17:47:27Z"),
		day:   parseDuration("12h7m1s"),
	},
	{
		place: "Kiritimati",
		lat:   1.8721,
		lon:   -157.4278,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-11T06:12:29+14:00"),
		set:   parseTime("2009-11-11T18:14:52+14:00"),
		day:   parseDuration("12h2m22s"),
	},
	{
		place: "Baker Island",
		lat:   0.1936,
		lon:   -176.4769,
		date:  parseTime("2009-11-10T23:00:00Z"),
		rise:  parseTime("2009-11-10T05:26:37-12:00"),
		set:   parseTime("2009-11-10T17:33:10-12:00"),
		day:   parseDuration("12h6m33s"),
	},
	// Errors
	{
		place: "North Pole",
		lat:   90.0,
		date:  parseTime("2009-11-10T23:00:00Z"),
		err:   ErrNoRise,
	},
	{
		place: "South Pole",
		lat:   -90.0,
		date:  parseTime("2009-11-10T23:00:00Z"),
		err:   ErrNoSet,
		day:   parseDuration("24h"),
	},
}

func TestRise(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.place, func(t *testing.T) {
			got, err := Rise(tc.date, tc.lat, tc.lon)
			if err != tc.err {
				t.Errorf("got %v; want %v", err, tc.err)
			}
			if got = got.In(tc.rise.Location()).Truncate(time.Second); !got.Equal(tc.rise) {
				t.Errorf("got %v; want %v", got, tc.rise)
			}
		})
	}
}

func TestSet(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.place, func(t *testing.T) {
			got, err := Set(tc.date, tc.lat, tc.lon)
			if err != tc.err {
				t.Errorf("got %v; want %v", err, tc.err)
			}
			if got = got.In(tc.set.Location()).Truncate(time.Second); !got.Equal(tc.set) {
				t.Errorf("got %v; want %v", got, tc.set)
			}
		})
	}
}

func TestDayDuration(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.place, func(t *testing.T) {
			got := Day(tc.date, tc.lat, tc.lon)
			if got = got.Truncate(time.Second); got != tc.day {
				t.Errorf("got %v; want %v", got, tc.day)
			}
		})
	}
}

func BenchmarkRise(b *testing.B) {
	for _, tc := range testCases {
		b.Run(tc.place, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Rise(tc.date, tc.lat, tc.lon)
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	for _, tc := range testCases {
		b.Run(tc.place, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Set(tc.date, tc.lat, tc.lon)
			}
		})
	}
}
