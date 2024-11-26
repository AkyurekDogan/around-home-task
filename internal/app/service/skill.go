/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/dto"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

// PartnerSkill interface provides partner skills methods
type PartnerSkill interface {
	Get(filter dto.Filter) (*dto.Skill, error)
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
func (s *partnerSkill) Get(filter dto.Filter) (*dto.Skill, error) {
	eFilter := s.toModel(filter)
	partnerSkill, err := s.dbPartnerSkill.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDTO(*partnerSkill)
	return &result, nil
}

func (s *partnerSkill) toDTO(pr model.Skill) dto.Skill {
	return pr.Skills
}

func (s *partnerSkill) toModel(p dto.Filter) model.Filter {
	return model.Filter{
		PartnerId: p.PartnerId,
	}
}
