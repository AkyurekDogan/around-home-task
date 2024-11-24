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

// PartnerRating represents the partner ratings data
type PartnerRating struct {
	PartnerId string
	Rating    int
}

// PartnerSkill represents the partner skills data
type PartnerSkill struct {
	PartnerId string
	Skills    []string
}

// Partners represents the partners data
type Partners []Partner
