/*
The Domain package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package domain

// Partner represents the partner data
type Partner struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Loc     Location       `json:"location"`
	Radius  Measure        `json:"radius"`
	PRating *PartnerRating `json:"partner_rating"`
	PSkills *PartnerSkill  `json:"partner_skill"`
}

type Rating struct {
	Value  float32 `json:"value"`
	Method string  `json:"method"`
}

// Rating represents the partner ratings data
type PartnerRating struct {
	Rating Rating `json:"ratings"`
}

// PartnerSkill represents the partner skills data
type PartnerSkill struct {
	Skill []string `json:"skills"`
}

// Partners represents the partners data
type Partners []Partner
