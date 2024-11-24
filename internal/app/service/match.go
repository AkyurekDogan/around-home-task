/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

// Match interface provides methods for matching partner-customer
type Match interface {
	Find(filter domain.MatchFilter) (*domain.MatchListResponse, error)
}

type match struct {
	dbMatch repository.Match
}

// NewMatch creates a new match service for matching
func NewMatch(db repository.Match) Match {
	return &match{
		dbMatch: db,
	}
}

// Get returns the matching records by customer filter details
func (s *match) Find(filter domain.MatchFilter) (*domain.MatchListResponse, error) {
	flt := s.toEntityFilter(filter)
	data, err := s.dbMatch.Get(flt)
	if err != nil {
		return nil, err
	}
	domData := s.toDomainMatchList(data)
	result := domain.MatchListResponse{
		Filter: filter,
		Data:   domData,
	}
	return &result, nil
}

func (s *match) toEntityFilter(filter domain.MatchFilter) entity.MatchFilter {
	return entity.MatchFilter{
		MaterialType: filter.MaterialType,
		Loc: entity.Location{
			Lat:  filter.Loc.Lat,
			Long: filter.Loc.Long,
		},
	}
}

func (s *match) toDomainMatchList(eml entity.MatchList) domain.MatchList {
	domML := make(domain.MatchList, 0, len(eml))
	for _, v := range eml {
		domML = append(domML, s.toDomainMatch(v))
	}
	return domML
}

func (s *match) toDomainMatch(m entity.Match) domain.Match {
	return domain.Match{
		PartnerId: m.PartnerId,
		Name:      m.Name,
		Loc: domain.Location{
			Lat:  m.Loc.Lat,
			Long: m.Loc.Long,
		},
		Rank: m.Rank,
	}
}
