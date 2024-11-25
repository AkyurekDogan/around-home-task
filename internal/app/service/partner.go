/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"errors"

	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/entity"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
)

var (
	ErrNoPartner = errors.New("partner is not found with given parameters")
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
		if errors.Is(err, repository.ErrNoRows) {
			return nil, ErrNoPartner
		}
		return nil, err
	}
	result := s.toDomain(partner)
	err = s.getPartnerSkills(&result, filter)
	if err != nil {
		return nil, err
	}
	err = s.getPartnerRating(&result, filter)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *partner) getPartnerSkills(p *domain.Partner, f domain.Filter) error {
	skill, err := s.svcPartnerSkill.Get(f)
	if err != nil {
		if !errors.Is(err, repository.ErrNoRows) {
			return err
		}
	}
	p.Skills = skill

	return nil
}

func (s *partner) getPartnerRating(p *domain.Partner, f domain.Filter) error {
	rating, err := s.svcPartnerRating.Get(f)
	if err != nil {
		if !errors.Is(err, repository.ErrNoRows) {
			return err
		}
	}
	p.Rating = rating
	return nil
}
func (s *partner) toDomain(p *entity.Partner) domain.Partner {
	return domain.Partner{
		Id:   p.Id,
		Name: p.Name,
		Loc: domain.Location{
			Lat:  p.Loc.Lat,
			Long: p.Loc.Long,
		},
		Radius: domain.Measure{
			Value:  float32(p.Radius),
			Metric: metricDistanceKM,
		},
	}
}

func (s *partner) toEntityFilter(p domain.Filter) entity.Filter {
	return entity.Filter{
		PartnerId: p.PartnerId,
	}
}
