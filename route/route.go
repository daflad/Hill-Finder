package route

import (
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"time"
)

//Route A sequence of GPX points
type Route struct {
	ID       int64
	Name     string
	Date     string
	Dist     float64
	Ascent   float64
	AvgSpeed float64
	Data     GPX
	Sections []Section
	Hills    []Hill
}

func (r *Route) String() string {
	return fmt.Sprintf("%v -- %v\nDistance: %.2f\nAscent: %.2fm\nAvg Speed: %.2f\n"+
		"Number Hills: %d", r.Name, r.Date, metersTomiles(r.Dist), r.Ascent, r.AvgSpeed, len(r.Hills))
}

func metersTomiles(distance float64) float64 {
	return distance / 1000 / 8 * 5
}

//GPX parsing from XML
type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Track   Track    `xml:"trk"`
}

//Track parsing from XML
type Track struct {
	XMLName  xml.Name `xml:"trk"`
	Name     string   `xml:"name"`
	Segments Segment  `xml:"trkseg"`
}

//Segment parsing from XML
type Segment struct {
	XMLName   xml.Name   `xml:"trkseg"`
	Locations []Location `xml:"trkpt"`
}

//Location parsing from XML
type Location struct {
	XMLName           xml.Name  `xml:"trkpt"`
	Latitude          float64   `xml:"lat,attr"`
	Longitude         float64   `xml:"lon,attr"`
	Elevation         float64   `xml:"ele"`
	Time              time.Time `xml:"time"`
	Gradient          float64
	DistanceFromStart float64
}

//Open a gpx file & decode the XML
func (g *GPX) Open(filePath string) error {
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()
		err = xml.NewDecoder(file).Decode(g)
	}
	return err
}

//GetMetrics Calculate extra route metrics
//
// Ascent 		: total elevation gain throughout route
// Dist   		: total distance of route
// DistanceFromStart: the distance from the start to the current location
func (r *Route) GetMetrics() {
	r.Name = r.Data.Track.Name
	// Itterate list for data
	for i := 0; i < len(r.Data.Track.Segments.Locations); i++ {
		r.Ascent += r.GetElevGain(i)
		r.Dist += r.GetDistance(i)
		r.Data.Track.Segments.Locations[i].DistanceFromStart = r.Dist
	}
	// Speed = distance / time km/h
	locations := r.Data.Track.Segments.Locations
	t := locations[len(locations)-1].Time.Unix() - locations[0].Time.Unix()
	r.AvgSpeed = r.Dist / 1000 / (float64(t) / 60 / 60)
	r.Date = locations[0].Time.Format("2006-01-02")
}

//GetElevGain Calculate the difference in elevation if +ive
func (r *Route) GetElevGain(index int) float64 {
	// Don't want to go out of bounds!
	if index < 1 {
		return 0
	}
	// only return diff if gain in elevation found
	locations := r.Data.Track.Segments.Locations
	diff := locations[index].Elevation - locations[index-1].Elevation
	if diff > 0 {
		return diff
	}
	return 0
}

//GetDistance Calculate the distance between 2 Locations
func (r *Route) GetDistance(index int) float64 {
	// Don't want to go out of bounds!
	if index < 1 {
		return 0
	}
	locations := r.Data.Track.Segments.Locations
	dist := r.Distance3D(locations[index], locations[index-1])
	return dist
}

//Distance2D This uses the ‘haversine’ formula to calculate the great-circle distance between two points – that is,
// the shortest distance over the earth’s surface – giving an ‘as-the-crow-flies’ distance between the points
// (ignoring any hills, of course!).
func (r *Route) Distance2D(a, b Location) float64 {
	// Earth's radius in meters
	EarthsRadius := 6371000.0

	// Convert the diff in location, converted to radians
	dLat := r.DegToRad(b.Latitude - a.Latitude)
	dLon := r.DegToRad(b.Longitude - a.Longitude)

	// The square of half the chord length between Locations
	re := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(r.DegToRad(a.Latitude))*math.Cos(r.DegToRad(b.Latitude))*math.Sin(dLon/2)*math.Sin(dLon/2)

	// The angular distance in radians
	c := 2 * math.Atan2(math.Sqrt(re), math.Sqrt(1-re))

	return EarthsRadius * c
}

//Distance3D Calculate the distance between 2 locations given the elevation of each point
func (r *Route) Distance3D(a, b Location) float64 {
	planar := r.Distance2D(a, b)
	height := math.Abs(b.Elevation - a.Elevation)
	return math.Sqrt(math.Pow(planar, 2) + math.Pow(height, 2))
}

//DegToRad Convert from degrees to radians
func (r *Route) DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}
