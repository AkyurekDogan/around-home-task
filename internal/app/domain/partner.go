/*
The Domain package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package domain

// Partner represents the partner data
type Partner struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Loc     Location      `json:"location"`
	Radius  int           `json:"radius"`
	PRating PartnerRating `json:"rating"`
	PSkills PartnerSkill  `json:"skills"`
}

// Rating represents the partner ratings data
type PartnerRating struct {
	Rating int `json:"rating"`
}

// PartnerSkill represents the partner skills data
type PartnerSkill struct {
	Skill []string `json:"skills"`
}

// Partners represents the partners data
type Partners []Partner