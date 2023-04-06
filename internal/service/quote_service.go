package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/sesaquecruz/goexpert-client-server-api-lab/internal/model"
)

type QuoteService struct {
	url string
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		url: "https://economia.awesomeapi.com.br/json/last/",
	}
}

func (a *QuoteService) GetQuote(ctx context.Context, pair string) (*model.Quote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.url+pair, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("invalid pair")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pairs map[string]model.Quote
	err = json.Unmarshal(body, &pairs)
	if err != nil {
		return nil, err
	}

	for _, quote := range pairs {
		return &quote, nil
	}

	return nil, errors.New("quote not found")
}
