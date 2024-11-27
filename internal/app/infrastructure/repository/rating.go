/*
Repository package for data access
*/
package repository

import (
	"database/sql"

	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/model"
)

// Rating represents the repository access layer for partner ratings
type Rating interface {
	Get(f model.Filter) (*model.Rating, error)
}

type rating struct {
	driverRead drivers.Driver
}

// NewRating creates new database interface for partner rating
func NewRating(driverR drivers.Driver) Rating {
	return &rating{
		driverRead: driverR,
	}
}

// Get gets the partner rating data
func (u *rating) Get(f model.Filter) (*model.Rating, error) {
	db, err := u.driverRead.Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var result model.Rating
	// Execute a SELECT query
	err = db.QueryRow(`
		select
			p.partner_id,
			p.avg
		from public.rating p
		where partner_id=$1
	`, f.PartnerId).Scan(&result.PartnerId, &result.ValueAVG)
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
