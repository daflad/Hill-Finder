package main

import (
	"log"

	"github.com/daflad/Hill-Finder/route"
)

func main() {
	var r route.Route
	r.Data.Open("TestData/Mynydd_beris_Capel_idwal_home.gpx")
	r.GetMetrics()
	r.FindClimbs()
	for _, hill := range r.Hills {
		log.Println(hill.String())
	}
}
