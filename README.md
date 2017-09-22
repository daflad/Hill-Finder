# Hill-Finder
A program parsing GXP routes and isolating the hills

## What is a climb?

A climb is any piece of road which is greater than 100 meters
in length with an average grade of 3% or more.

grade = (vertical climb / horizontal distance) * 100

## Grade bands


*  3 -  4 %
*  5 -  7 %
*  8 -  9 %
* 10 - 15 %
* 15 +    %



## Classifications (UCI - Standard)


* 4th Category 		– 	climbs of  100 -  300 meters
* 3rd Category 		– 	climbs of  300 -  600 meters
* 2nd Category 		– 	climbs of  600 - 1100 meters
* 1st Category 		– 	climbs of 1100 - 1500 meters
* Hors Category (HC) 	– 	climbs of 1500 +      meters



## The distance factor

The distsance into the ride has a big impact on climb severity

Distance Factor

*   0 - 20 miles - 1.0
*  20 - 40 miles - 1.1
*  40 - 60 miles - 1.2
*  60 - 70 miles - 1.3
*  70 - 80 miles - 1.4
*  80 - 90 miles - 1.5
* 100 +		  - 1.6



## Calculating a category

Length of climb(m) * grade (%) * distance factor

* 4th Category 		> 	 8000
* 3rd Category 		> 	16000
* 2nd Category 		> 	32000
* 1st Category 		> 	64000
* Hors Category (HC) 	> 	80000
