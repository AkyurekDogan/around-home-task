/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

// PartnerRating interface provides partner rating methods
type PartnerRating interface {
	Get(filter domain.Filter) (*domain.PartnerRating, error)
}

type partnerRating struct {
	dbPartnerRating repository.PartnerRating
}

// NewPartnerRating creates a new partner rating service to access relevant operations
func NewPartnerRating(repoPartnerRating repository.PartnerRating) PartnerRating {
	return &partnerRating{
		dbPartnerRating: repoPartnerRating,
	}
}

// Get returns the relavent partner rating data by filter
func (s *partnerRating) Get(filter domain.Filter) (*domain.PartnerRating, error) {
	eFilter := s.toEntityFilter(filter)
	partnerRating, err := s.dbPartnerRating.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDomain(*partnerRating)
	return &result, nil
}

func (s *partnerRating) toDomain(pr entity.PartnerRating) domain.PartnerRating {
	return domain.PartnerRating{
		Rating: pr.Rating,
	}
}

func (s *partnerRating) toEntityFilter(p domain.Filter) entity.Filter {
	return entity.Filter{
		PartnerId: p.PartnerId,
	}
}