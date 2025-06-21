package services_test

import (
	"backend/core/models"
	"backend/core/repositories"
	"backend/core/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetQuotes(t *testing.T) {
	type test struct {
		Name string
		Mock struct {
			GetQuotes struct {
				Output []models.QuoteModel
				Error  error
			}
		}
		Output models.ResponseModel
	}
	cases := []test{
		{
			Name: "get quotes success",
			Mock: struct {
				GetQuotes struct {
					Output []models.QuoteModel
					Error  error
				}
			}{
				GetQuotes: struct {
					Output []models.QuoteModel
					Error  error
				}{
					Output: []models.QuoteModel{},
					Error:  nil,
				},
			},
			Output: models.ResponseModel{
				Status:  true,
				Code:    200,
				Message: "get quotes success",
				Result:  []models.QuoteModel{},
			},
		},
		{
			Name: "get quotes error",
			Mock: struct {
				GetQuotes struct {
					Output []models.QuoteModel
					Error  error
				}
			}{
				GetQuotes: struct {
					Output []models.QuoteModel
					Error  error
				}{
					Output: nil,
					Error:  errors.New("error"),
				},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "error",
				Result:  nil,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			quoteRepo := repositories.NewQuoteRepositoryMock()
			quoteRepo.On("GetQuotes").Return(c.Mock.GetQuotes.Output, c.Mock.GetQuotes.Error)

			quoteService := services.NewQuoteService(quoteRepo)
			result := quoteService.GetQuotes()

			assert.Equal(t, c.Output, result)
		})
	}
}
