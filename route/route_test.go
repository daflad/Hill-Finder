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
	if elev != 287.0 {
		t.Error("Expected 287.0, got ", elev)
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

func TestElevationGain(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/SimpleRoute.gpx")
	gain := r.GetElevGain(1)
	if gain != 0.5 {
		t.Error("Expected 0.5, got ", gain)
	}
	noGain := r.GetElevGain(0)
	if noGain != 0 {
		t.Error("Expected 0, got ", noGain)
	}
}

func TestGetDistance(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/SimpleRoute.gpx")
	dist := r.GetDistance(1)
	if dist != 3.792887554216133 {
		t.Error("Expected 3.792887554216133, got ", dist)
	}
	noDist := r.GetDistance(0)
	if noDist != 0 {
		t.Error("Expected 0, got ", noDist)
	}
}
