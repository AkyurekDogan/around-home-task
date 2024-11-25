/*
The Entity package for database access and compatible for relevant database
*/
package entity

// Partner represents the partner data
type Partner struct {
	Id     string
	Name   string
	Loc    Location
	Radius int
}

type Rating struct {
	Value  int
	Method string
}

// PartnerRating represents the partner ratings data
type PartnerRating struct {
	PartnerId string
	Rating    Rating
}

// PartnerSkill represents the partner skills data
type PartnerSkill struct {
	PartnerId string
	Skills    []string
}

// Partners represents the partners data
type Partners []Partner
