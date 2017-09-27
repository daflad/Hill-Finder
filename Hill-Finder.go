package main

import (
	"log"

	"github.com/daflad/Hill-Finder/route"
)

func main() {
	var r route.Route
	r.Data.Open("TestData/Ogwen, beris, tragrth, Ogwen .gpx")
	r.GetMetrics()
	r.FindClimbs()
	for _, hill := range r.Hills {
		log.Println(hill.Category())
	}
}
