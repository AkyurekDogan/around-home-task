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
	Get(filter domain.Filter) (*domain.Skill, error)
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
func (s *partnerSkill) Get(filter domain.Filter) (*domain.Skill, error) {
	eFilter := s.toEntityFilter(filter)
	partnerSkill, err := s.dbPartnerSkill.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDomain(*partnerSkill)
	return &result, nil
}

func (s *partnerSkill) toDomain(pr entity.Skill) domain.Skill {
	return pr.Skills
}

func (s *partnerSkill) toEntityFilter(p domain.Filter) entity.Filter {
	return entity.Filter{
		PartnerId: p.PartnerId,
	}
}
