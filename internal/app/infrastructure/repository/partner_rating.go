/*
Repository package for data access
*/
package repository

import (
	"database/sql"

	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
)

// PartnerRating represents the repository access layer for partner ratings
type PartnerRating interface {
	Get(f entity.Filter) (*entity.PartnerRating, error)
}

type partnerRating struct {
	driverRead drivers.Driver
}

// NewPartnerSkill creates new database interface for PartnerSkill
func NewPartnerRating(driverR drivers.Driver) PartnerRating {
	return &partnerRating{
		driverRead: driverR,
	}
}

// Get gets the partner skill data
func (u *partnerRating) Get(f entity.Filter) (*entity.PartnerRating, error) {
	db, err := u.driverRead.Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var result entity.PartnerRating
	// Execute a SELECT query
	err = db.QueryRow(`
		select
			p.partner_id,
			p.rating,
		from public.partner_rating p
		where partner_id=$s
	`, f.PartnerId).Scan(&result.PartnerId, &result.Rating)
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