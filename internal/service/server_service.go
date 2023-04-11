package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sesaquecruz/goexpert-client-server-api-lab/internal/database/repository"
)

type ServerService struct {
	quoteRepository *repository.QuoteRepository
	quoteService    *QuoteService
	ApiTimeout      time.Duration
	DbTimeout       time.Duration
}

func NewServerService(
	qr *repository.QuoteRepository,
	qs *QuoteService,
	apiTimeout time.Duration,
	dbTimeout time.Duration,
) *ServerService {
	return &ServerService{
		quoteRepository: qr,
		quoteService:    qs,
		ApiTimeout:      apiTimeout,
		DbTimeout:       dbTimeout,
	}
}

func (s *ServerService) UsdBrlHandler(w http.ResponseWriter, r *http.Request) {
	// Get quote from API
	ctx, cancel := context.WithTimeout(r.Context(), s.ApiTimeout)
	defer cancel()

	quote, err := s.quoteService.GetQuote(ctx, "USD-BRL")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// Save quote on DB
	ctx, cancel = context.WithTimeout(r.Context(), s.DbTimeout)
	defer cancel()
	err = s.quoteRepository.SaveQuote(ctx, quote)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return bid field
	json.NewEncoder(w).Encode(&struct {
		Bid float64 `json:"bid,string"`
	}{
		Bid: quote.Bid,
	})
}
