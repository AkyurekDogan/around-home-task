/*
Repository package for data access
*/
package repository

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
)

// Match represents the usage database interface
type Match interface {
	Get(filter entity.MatchFilter) (entity.MatchList, error)
}

type match struct {
	driverRead drivers.Driver
}

// NewMatch creates new database interface to access database records
func NewMatch(driverR drivers.Driver) Match {
	return &match{
		driverRead: driverR,
	}
}

// Get gets the records regarding to the filter list
func (u *match) Get(filter entity.MatchFilter) (entity.MatchList, error) {
	db, err := u.driverRead.Init()
	if err != nil {
		return nil, err
	}
	// Execute a SELECT query
	rows, err := db.Query(`
		select
			u.instance_id,
			u.sku,
			u.start_time,
			u.end_time,
			u.unit_size
		from public.usage u
		where u.start_time >= $1
			and u.end_time <= $2
		order by u.instance_id
		limit $3 offset $4
	`, filter.Loc.Lat, filter.Loc.Long, filter.MaterialType, filter.MaterialType) // #TODO fix it
	if err != nil {
		return nil, err
	}
	defer db.Close()
	defer rows.Close()

	var result entity.MatchList
	// Iterate through the result set
	for rows.Next() {
		item := entity.Match{}
		err := rows.Scan(&item.PartnerId, &item.Name, &item.Distance, &item.Loc, &item.Rank)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}
