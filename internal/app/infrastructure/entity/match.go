/*
The Entity package for database access and compatible for relevant database
*/
package entity

// Location struct for coordinates
type Location struct {
	Lat  float64
	Long float64
}

// Match struct for cost data
type Match struct {
	PartnerId string
	Name      string
	Loc       Location
	Rank      int
	Distance  int
}

// MatchList represents the match data collection as list
type MatchList []Match
