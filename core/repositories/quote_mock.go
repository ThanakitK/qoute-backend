package repositories

import (
	"backend/core/models"

	"github.com/stretchr/testify/mock"
)

type quoteRepoMock struct {
	mock.Mock
}

func NewQuoteRepositoryMock() *quoteRepoMock {
	return &quoteRepoMock{}
}

func (m *quoteRepoMock) GetQuotes() (result []models.QuoteModel, err error) {
	args := m.Called()
	return args.Get(0).([]models.QuoteModel), args.Error(1)
}

func (m *quoteRepoMock) GetQuote(id string) (result models.QuoteModel, err error) {
	args := m.Called(id)
	return args.Get(0).(models.QuoteModel), args.Error(1)
}

func (m *quoteRepoMock) CreateQuote(payload models.CreateQuoteModel) (result models.QuoteModel, err error) {
	args := m.Called(payload)
	return args.Get(0).(models.QuoteModel), args.Error(1)
}

func (m *quoteRepoMock) UpdateQuote(id string, payload models.UpdateQuoteModel) (result models.QuoteModel, err error) {
	args := m.Called(id, payload)
	return args.Get(0).(models.QuoteModel), args.Error(1)
}

func (m *quoteRepoMock) DeleteQuote(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
