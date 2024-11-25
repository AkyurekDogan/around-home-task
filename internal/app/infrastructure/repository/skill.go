/*
Repository package for data access
*/
package repository

import (
	"database/sql"

	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
	"github.com/lib/pq"
)

// PartnerSkill represents the repository access layer for partner
type PartnerSkill interface {
	Get(f entity.Filter) (*entity.Skill, error)
}

type partnerSkill struct {
	driverRead drivers.Driver
}

// NewPartnerSkill creates new database interface for PartnerSkill
func NewPartnerSkill(driverR drivers.Driver) PartnerSkill {
	return &partnerSkill{
		driverRead: driverR,
	}
}

// Get gets the partner skill data
func (u *partnerSkill) Get(f entity.Filter) (*entity.Skill, error) {
	db, err := u.driverRead.Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var result entity.Skill
	// Execute a SELECT query
	err = db.QueryRow(`
		select
			p.partner_id,
			p.craftsmanship_tags
		from public.skill p
		where partner_id=$1
	`, f.PartnerId).Scan(&result.PartnerId, pq.Array(&result.Skills))
	if err != nil {
		// Check for no rows found or other errors
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		} else {
			return nil, err
		}
	}
	return &result, nil
}
