/*
The Domain package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package domain

// Location struct for coordinates
type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// Match struct for cost data
type Match struct {
	PartnerId string   `json:"partner_id"`
	Name      string   `json:"name"`
	Loc       Location `json:"location"`
	Rank      int      `json:"rank"`
}

// MatchList represents the response struct that for success response
type MatchList []Match

// MatchList represents the response struct that for success response
type MatchListResponse struct {
	Filter MatchFilter `json:"filter"`
	Data   MatchList   `json:"data"`
}