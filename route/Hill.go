package route

//Hill A hill found in a route
type Hill struct {
	ID           int
	Name         string
	Start        Section
	End          Section
	Sections     []Section
	AverageGrade float64
	MaxGrade     float64
	MinGrade     float64
}

//Section A section of a hill
type Section struct {
	Start Location
	End   Location
	Grade float64
}

//FindClimbs Traverse the list of Locations to split route into 100m sections
func (r *Route) FindClimbs() {
	locations := r.Data.Track.Segments.Locations
	var sec Section
	// Markers to help logic in sections:
	// s = started
	// e = finished
	s := false
	e := false
	for i := 0; i < len(locations); i++ {
		if !s {
			sec.Start = locations[i]
			s = true
		}
		if !e {
			if locations[i].DistanceFromStart-sec.Start.DistanceFromStart > 299 {
				sec.End = locations[i]
				sec.Grade = ((sec.End.Elevation - sec.Start.Elevation) / (sec.End.DistanceFromStart - sec.Start.DistanceFromStart)) * 100
				s = false
				e = true
			}
		}
		if e {
			// Append section to list
			r.Sections = append(r.Sections, sec)
			e = false
		}
	}

	// Markers to help logic in sections:
	// s = started
	// e = finished
	// s = false
	e = false
	for i := 0; e || i < len(r.Sections); i++ {

		if r.Sections[i].Grade > 3 {
			var hill Hill
			hill.Start = r.Sections[i]
			//log.Printf("Start Location of Hill: %v,%v grade: %v\n", hill.Start.Start.Latitude,
			//	hill.Start.Start.Longitude, r.Sections[i].Grade)
			hill.AverageGrade = hill.Start.Grade
			for j := i; j < len(r.Sections); j++ {
				if (hill.AverageGrade+r.Sections[j].Grade)/float64(len(hill.Sections)) < 3 {
					spare := len(r.Sections) - j
					if spare > 3 {
						spare = 3
					}
					tempAVG := hill.AverageGrade
					for k := j; k < j+spare; k++ {
						tempAVG += r.Sections[k].Grade
					}
					if tempAVG < 3 {
						hill.End = r.Sections[j-1]
						//log.Printf("End Location of Hill: %v,%v grade: %v\n", hill.End.End.Latitude,
						//	hill.End.End.Longitude, r.Sections[i].Grade)
						i = j
						r.Hills = append(r.Hills, hill)
						// fmt.Printf("Hill found, length: %4d m with avg Grade : %.2f%%", len(hill.Sections)*300, hill.AverageGrade)
						break
					}
				} else {
					hill.Sections = append(hill.Sections, r.Sections[j])
					hill.AverageGrade = (hill.AverageGrade + r.Sections[j].Grade) / 2
					// fmt.Printf("Hill Section Added: %4d grade: %.2f%%\n", j, r.Sections[j].Grade)
				}
			}

		}
	}
	// log.Printf("# Hills found: %v", len(r.Hills))
	// var start, end Location
	// var hill Hill
}
