/*
Repository package for data access
*/
package repository

import (
	"database/sql"
	"errors"

	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/model"
)

var (
	ErrNoRows = errors.New("no record found with the given id")
)

// Partner represents the repository access layer for partner
type Partner interface {
	Get(f model.Filter) (*model.Partner, error)
}

type partner struct {
	driverRead drivers.Driver
}

// NewPartner creates new database interface for Partner
func NewPartner(driverR drivers.Driver) Partner {
	return &partner{
		driverRead: driverR,
	}
}

// Get gets the partner data
func (u *partner) Get(f model.Filter) (*model.Partner, error) {
	db, err := u.driverRead.Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var result model.Partner
	// Execute a SELECT query
	err = db.QueryRow(`
		select
			p.id,
			p.name,
			ST_X(location::geometry) as lat,
			ST_Y(location::geometry) as long,
			p.radius
		from public.partner p
		where id=$1
	`, f.PartnerId).Scan(&result.Id, &result.Name, &result.Loc.Lat, &result.Loc.Long, &result.Radius)
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
