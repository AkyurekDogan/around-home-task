/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

// PartnerSkill interface provides partner skills methods
type PartnerSkill interface {
	Get(filter domain.Filter) (*domain.PartnerSkill, error)
}

type partnerSkill struct {
	dbPartnerSkill repository.PartnerSkill
}

// NewPartnerSkill creates a new partner skills service to access relevant operations
func NewPartnerSkill(repoPartnerSkill repository.PartnerSkill) PartnerSkill {
	return &partnerSkill{
		dbPartnerSkill: repoPartnerSkill,
	}
}

// Get returns the relavent partner rating data by filter
func (s *partnerSkill) Get(filter domain.Filter) (*domain.PartnerSkill, error) {
	eFilter := s.toEntityFilter(filter)
	partnerRating, err := s.dbPartnerSkill.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDomain(*partnerRating)
	return &result, nil
}

func (s *partnerSkill) toDomain(pr entity.PartnerSkill) domain.PartnerSkill {
	return domain.PartnerSkill{
		Skill: pr.Skills,
	}
}

func (s *partnerSkill) toEntityFilter(p domain.Filter) entity.Filter {
	return entity.Filter{
		PartnerId: p.PartnerId,
	}
}
