package route

import "testing"

func TestOpen(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/SimpleRoute.gpx")
	v := len(r.Data.Track.Segments.Locations)
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
	fp := r.Data.Track.Segments.Locations[0]
	firstTime := fp.Time.UTC().String()
	if firstTime != "2016-08-02 17:13:54 +0000 UTC" {
		t.Error("Expected 2016-08-02 17:13:54 +0000 UTC, got ", firstTime)
	}
	elev := fp.Elevation
	if elev != 287.5 {
		t.Error("Expected 287.5, got ", elev)
	}
	lat := fp.Latitude
	if lat != 53.1362830 {
		t.Error("Expected 53.1362830, got ", lat)
	}
	lon := fp.Longitude
	if lon != -4.1234530 {
		t.Error("Expected -4.1234530, got ", lon)
	}
}
