package main

import "daflad/Hill-Finder/route"

func main() {
	var r route.Route
	r.Data.Open("TestData/Evening ride to Pen-y-pass and back.gpx")
	r.GetMetrics()
	r.FindClimbs()
}
