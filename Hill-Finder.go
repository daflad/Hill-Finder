package main

import "github.com/daflad/Hill-Finder/route"

func main() {
	var r route.Route
	r.Data.Open("TestData/Ogwen, beris, tragrth, Ogwen .gpx")
	r.GetMetrics()
	r.FindClimbs()
}
