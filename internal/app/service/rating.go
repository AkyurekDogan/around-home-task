/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/dto"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

// PartnerRating interface provides partner rating methods
type PartnerRating interface {
	Get(filter dto.Filter) (*dto.Rating, error)
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
func (s *partnerRating) Get(filter dto.Filter) (*dto.Rating, error) {
	eFilter := s.toModel(filter)
	partnerRating, err := s.dbPartnerRating.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDTO(*partnerRating)
	return &result, nil
}

func (s *partnerRating) toDTO(pr model.Rating) dto.Rating {
	return dto.Rating{
		ValueAVG: pr.ValueAVG,
	}
}

func (s *partnerRating) toModel(p dto.Filter) model.Filter {
	return model.Filter{
		PartnerId: p.PartnerId,
	}
}
