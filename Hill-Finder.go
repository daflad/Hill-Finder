package main

import (
	"fmt"

	"github.com/daflad/Hill-Finder/route"
)

func main() {
	var r route.Route
	r.Data.Open("TestData/Mynydd_beris_Capel_idwal_home.gpx")
	r.GetMetrics()
	r.FindClimbs()
	fmt.Println(r.String())
	for _, hill := range r.Hills {
		fmt.Println(hill.String())
	}
}
