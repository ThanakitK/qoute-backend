package services

import (
	"backend/core/models"
	"backend/core/repositories"
	"time"

	"github.com/google/uuid"
)

type QuoteService interface {
	GetQuotes() (result models.ResponseModel)

	CreateQuote(quote string) (result models.ResponseModel)

	UpdateQuote(id string, quote string, vote int) (result models.ResponseModel)

	DeleteQuote(id string) (result models.ResponseModel)
}
type QuoteSrv struct {
	quoteRepo repositories.QuoteRepository
}

func NewQuoteService(quoteRepo repositories.QuoteRepository) QuoteService {
	return &QuoteSrv{
		quoteRepo: quoteRepo,
	}
}

func (s *QuoteSrv) GetQuotes() (result models.ResponseModel) {
	res, err := s.quoteRepo.GetQuotes()
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}
	return models.ResponseModel{
		Status:  true,
		Code:    200,
		Message: "get quotes success",
		Result:  res,
	}
}

func (s *QuoteSrv) CreateQuote(quote string) (result models.ResponseModel) {
	payload := models.RepoCreateQuoteModel{
		ID:         uuid.New().String(),
		Quote:      quote,
		Vote:       0,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}
	res, err := s.quoteRepo.CreateQuote(payload)
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}
	return models.ResponseModel{
		Status:  true,
		Code:    201,
		Message: "create quote success",
		Result:  res,
	}
}

func (s *QuoteSrv) UpdateQuote(id string, quote string, vote int) (result models.ResponseModel) {
	payload := models.RepoUpdateQuoteModel{
		Quote:      quote,
		Vote:       vote,
		UpdateDate: time.Now(),
	}
	res, err := s.quoteRepo.UpdateQuote(id, payload)
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}

	return models.ResponseModel{
		Status:  true,
		Code:    200,
		Message: "update quote success",
		Result:  res,
	}
}

func (s *QuoteSrv) DeleteQuote(id string) (result models.ResponseModel) {
	res, err := s.quoteRepo.GetQuote(id)
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}
	if res.Vote > 0 {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "cannot delete quote with vote > 0",
			Result:  nil,
		}
	}
	err = s.quoteRepo.DeleteQuote(id)
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}
	return models.ResponseModel{
		Status:  true,
		Code:    200,
		Message: "delete quote success",
		Result:  nil,
	}
}
