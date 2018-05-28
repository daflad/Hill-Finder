package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/daflad/Hill-Finder/route"
)

func main() {
	if len(os.Args) > 1 {
		pathToGPX := os.Args[1]
		files, err := ioutil.ReadDir(pathToGPX)
		if err != nil {
			log.Println(err)
		} else {
			for _, file := range files {
				if strings.Contains(file.Name(), "Ride") {
					fp := pathToGPX + "/" + file.Name()
					if pathToGPX[len(pathToGPX)-1] == '/' {
						fp = pathToGPX + file.Name()
					}
					var r route.Route
					r.Data.Open(fp)
					r.GetMetrics()
					r.FindClimbs()
					if len(r.Hills) > 0 {
						fmt.Println(r.String())
						for _, hill := range r.Hills {
							fmt.Println(hill.String())
						}
					}
					break
				}
			}
		}
	}
}
