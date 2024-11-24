/*
The Handler package to manage the request-response pipeline handlers
*/
package handler

import (
	"net/http"

	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/service"
)

// Match represents the match handler
type Match interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type match struct {
	base
	matchService service.Match
}

// NewMatch creates new match handler
func NewMatch(ms service.Match) Match {
	return &match{
		matchService: ms,
	}
}

// Get returns the match results by filter
func (s *match) Get(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	filter, err := domain.NewMatchFilter(r)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid query parameters", err)
		return
	}
	// Validation check
	if err := filter.CheckFilter(); err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid or insufficient input", err)
		return
	}
	result, err := s.matchService.Find(*filter)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusInternalServerError, "internal server error", err)
		return
	}
	s.WriteSuccessRespone(w, http.StatusOK, result)
}
