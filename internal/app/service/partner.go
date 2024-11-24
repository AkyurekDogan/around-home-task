/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

// Partner interface provides price methods
type Partner interface {
	Get(filter domain.Filter) (*domain.Partner, error)
}

type partner struct {
	dbPartner        repository.Partner
	svcPartnerSkill  PartnerSkill
	svcPartnerRating PartnerRating
}

// NewPriceService creates a new price service with following operations: get
func NewPriceService(repoPartner repository.Partner,
	partnerSkillService PartnerSkill,
	partnerRatingService PartnerRating,
) Partner {
	return &partner{
		dbPartner:        repoPartner,
		svcPartnerSkill:  partnerSkillService,
		svcPartnerRating: partnerRatingService,
	}
}

// Get returns the relavent partner data
func (s *partner) Get(filter domain.Filter) (*domain.Partner, error) {
	eFilter := s.toEntityFilter(filter)
	partner, err := s.dbPartner.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDomain(*partner)
	skill, err := s.svcPartnerSkill.Get(filter)
	if err != nil {
		return nil, err
	}
	result.PSkills = *skill
	rating, err := s.svcPartnerRating.Get(filter)
	if err != nil {
		return nil, err
	}
	result.PRating = *rating
	return &result, nil
}

func (s *partner) toDomain(p entity.Partner) domain.Partner {
	return domain.Partner{
		Id:   p.Id,
		Name: p.Name,
		Loc: domain.Location{
			Lat:  p.Loc.Lat,
			Long: p.Loc.Long,
		},
		Radius: p.Radius,
	}
}

func (s *partner) toEntityFilter(p domain.Filter) entity.Filter {
	return entity.Filter{
		PartnerId: p.PartnerId,
	}
}
