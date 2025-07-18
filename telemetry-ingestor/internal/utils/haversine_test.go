package utils

import (
	"math"
	"testing"
)

func TestHaversine(t *testing.T) {
	margin := 0.1

	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Same location",
			args: args{lat1: 0, lon1: 0, lat2: 0, lon2: 0},
			want: 0,
		},
		{
			name: "New York to London",
			args: args{lat1: 40.7128, lon1: -74.0060, lat2: 51.5074, lon2: -0.1278},
			want: 5570222.18,
		},
		{
			name: "Sydney to Tokyo",
			args: args{lat1: -33.8688, lon1: 151.2093, lat2: 35.6895, lon2: 139.6917},
			want: 7826615.05,
		},
		{
			name: "Pune to Mumbai",
			args: args{lat1: 18.5204, lon1: 73.8567, lat2: 18.9582, lon2: 72.8321},
			want: 118364.64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Haversine(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2); math.Abs(got-tt.want) > margin {
				// if got := Haversine(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2); got != tt.want {
				t.Errorf("Haversine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHaversine(b *testing.B) {
	lat1, lon1 := 40.7128, -74.0060
	lat2, lon2 := 51.5074, -0.1278

	for i := 0; i < b.N; i++ {
		_ = Haversine(lat1, lon1, lat2, lon2)
	}
}
