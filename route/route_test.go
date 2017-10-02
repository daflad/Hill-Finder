package route

import "testing"

func TestOpen(t *testing.T) {
	var r Route
	//Testing invalid path
	err := r.Data.Open("fake/file/path")
	if err.Error() != "open fake/file/path: no such file or directory" {
		t.Error("Expected open fake/file/path: no such file or directory, got ", err.Error())
	}
	//Testing invalid gpx format
	err = r.Data.Open("../TestData/InvalidRoute.gpx")
	if err.Error() != "expected element type <gpx> but have <gpxs>" {
		t.Error("Expected expected element type <gpx> but have <gpxs>, got ", err.Error())
	}
	//Testing valid open
	r.Data.Open("../TestData/SimpleRoute.gpx")
	v := len(r.Data.Track.Segments.Locations)
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
	fp := r.Data.Track.Segments.Locations[0]
	firstTime := fp.Time.UTC().String()
	//Testing parsing
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
	noGainValid := r.GetElevGain(2)
	if noGainValid != 0 {
		t.Error("Expected 0, got ", noGainValid)
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

func TestGetMetrics(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/SimpleRoute.gpx")
	r.GetMetrics()
	if r.Ascent != 0.5 {
		t.Error("Expected 0.5, got ", r.Ascent)
	}
	if r.Dist != 3.792887554216133 {
		t.Error("Expected 3.792887554216133, got ", r.Dist)
	}
	if r.AvgSpeed != 2.2757325325296796 {
		t.Error("Expected 2.2757325325296796, got ", r.AvgSpeed)
	}
	expected := `Test Route -- 2016-08-02
Distance: 0.00
Ascent: 0.00
Avg Speed: 2.28
Number Hills: 0`
	if expected != r.String() {
		t.Error("Expected different string, got ", r.String())
	}
}

func TestFindNoClimbs(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/SimpleRoute.gpx")
	r.FindClimbs()
	nHills := len(r.Hills)
	if nHills != 0 {
		t.Error("Expected 0, got ", nHills)
	}
}

func TestFindOneClimbs(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/Ogwen, beris, tragrth, Ogwen .gpx")
	r.GetMetrics()
	r.FindClimbs()
	sHills := len(r.Hills)
	if sHills != 1 {
		t.Error("Expected 1, got ", sHills)
	}
	expected := "DFS:: 8.30,\tLEN:: 2.31,\tAGR:: 9.12%,\tMGR:: 12.93%,\tCAT:: 4th"
	if expected != r.Hills[0].String() {
		t.Error("Expected", expected, ", got ", r.Hills[0].String())
	}
}

func TestCategories(t *testing.T) {
	var r Route
	r.Data.Open("../TestData/Ogwen, beris, tragrth, Ogwen .gpx")
	r.GetMetrics()
	r.FindClimbs()
	r.Hills[0].Start.Start.DistanceFromStart = 80000
	expected := "3rd"
	if expected != r.Hills[0].Category() {
		t.Error("Expected 3rd, got ", r.Hills[0].Category())
	}
	r.Hills[0].Start.Start.DistanceFromStart = 100000
	if expected != r.Hills[0].Category() {
		t.Error("Expected 3rd, got ", r.Hills[0].Category())
	}
	r.Hills[0].Start.Start.DistanceFromStart = 120000
	if expected != r.Hills[0].Category() {
		t.Error("Expected 3rd, got ", r.Hills[0].Category())
	}
	r.Hills[0].Start.Start.DistanceFromStart = 140000
	r.Hills[0].AverageGrade = 30
	expected = "1st"
	if expected != r.Hills[0].Category() {
		t.Error("Expected 3rd, got ", r.Hills[0].Category())
	}
	r.Hills[0].Start.Start.DistanceFromStart = 150000
	r.Hills[0].AverageGrade = 15
	expected = "2nd"
	if expected != r.Hills[0].Category() {
		t.Error("Expected 2nd, got ", r.Hills[0].Category())
	}
	r.Hills[0].Start.Start.DistanceFromStart = 190000
	r.Hills[0].AverageGrade = 50
	expected = "HC"
	if expected != r.Hills[0].Category() {
		t.Error("Expected HC, got ", r.Hills[0].Category())
	}
}

func TestMetersToMiles(t *testing.T) {
	v := metersTomiles(8000)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}
