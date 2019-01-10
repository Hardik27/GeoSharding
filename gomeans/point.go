package gomeans

import "math"

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)
//Point struct is a simple coordinate
type Point struct {
	Lat float64
	Lon float64
	Ip string
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

//Distance function calculates distance between two points in the cartesian plan
func (p Point) Distance(q Point) float64 {
	
	//euclidean distance
	// first := math.Pow(float64(p2.X-p.X), 2)
	// second := math.Pow(float64(p2.Y-p.Y), 2)
	// return math.Sqrt(first + second)


	//haversine distance
	//Distance calculates the shortest path between two coordinates on the surface
	// of the Earth. This function returns distance in kilometers.
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	//mi = c * earthRadiusMi
	km := c * earthRaidusKm

	return km
}
