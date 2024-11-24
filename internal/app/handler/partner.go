/*
The Handler package to manage the request-response pipeline handlers
*/
package handler

import (
	"net/http"

	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/service"
)

// Partner represents the partner handler
type Partner interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type partner struct {
	base
	partnerService service.Partner
}

// NewPartner returns the new partner service
func NewPartner(ps service.Partner) Partner {
	return &partner{
		partnerService: ps,
	}
}

// Get returns the cost data
func (s *partner) Get(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	filter, err := domain.NewFilter(r)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid query parameters", err)
		return
	}
	// Validation check
	if err := filter.CheckFilter(); err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid or insufficient input", err)
		return
	}
	result, err := s.partnerService.Get(*filter)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusInternalServerError, "internal server error", err)
		return
	}
	s.WriteSuccessRespone(w, http.StatusOK, result)
}
