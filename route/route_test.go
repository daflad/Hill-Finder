package route

import (
	"testing"
)

func TestOpen(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/SimpleRoute.gpx")
	v := len(r.Data.Track.Segments.Locations)
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}
