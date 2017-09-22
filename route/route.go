package route

import (
	"encoding/xml"
	"log"
	"os"
)

//Route A sequence of GPX points
type Route struct {
	ID   int64
	Name string
	Date string
	Data GPX
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
	XMLName   xml.Name `xml:"trkpt"`
	Latitude  float64  `xml:"lat,attr"`
	Longitude float64  `xml:"lon,attr"`
	Elevation float64  `xml:"ele"`
	Time      string   `xml:"time"`
	Gradient  float64
}

//Open a gpx file & decode the XML
func (g *GPX) Open(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := xml.NewDecoder(file).Decode(g); err != nil {
		log.Fatal(err)
	}
}
