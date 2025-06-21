package services_test

import (
	"backend/core/models"
	"backend/core/repositories"
	"backend/core/services"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func Test_CreateQuote(t *testing.T) {
	type test struct {
		Name  string
		Input struct {
			Quote string
		}
		Mock struct {
			CreateQuote struct {
				Input  models.CreateQuoteModel
				Output models.QuoteModel
				Error  error
			}
		}
		Output models.ResponseModel
	}
	id := uuid.New().String()
	date := time.Now()
	cases := []test{
		{
			Name: "create quote success",
			Input: struct {
				Quote string
			}{
				Quote: "quote",
			},
			Mock: struct {
				CreateQuote struct {
					Input  models.CreateQuoteModel
					Output models.QuoteModel
					Error  error
				}
			}{
				CreateQuote: struct {
					Input  models.CreateQuoteModel
					Output models.QuoteModel
					Error  error
				}{
					Input: models.CreateQuoteModel{
						ID:         id,
						Quote:      "quote",
						CreateDate: date,
						UpdateDate: date,
					},
					Output: models.QuoteModel{
						ID:         id,
						Quote:      "quote",
						CreateDate: date,
						UpdateDate: date,
					},
					Error: nil,
				},
			},
			Output: models.ResponseModel{
				Status:  true,
				Code:    201,
				Message: "create quote success",
				Result: models.QuoteModel{
					ID:         id,
					Quote:      "quote",
					CreateDate: date,
					UpdateDate: date,
				},
			},
		},
		{
			Name: "quote not found",
			Input: struct {
				Quote string
			}{
				Quote: "",
			},
			Mock: struct {
				CreateQuote struct {
					Input  models.CreateQuoteModel
					Output models.QuoteModel
					Error  error
				}
			}{
				CreateQuote: struct {
					Input  models.CreateQuoteModel
					Output models.QuoteModel
					Error  error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "quote not found",
				Result:  nil,
			},
		},
		{
			Name: "create quote error",
			Input: struct {
				Quote string
			}{
				Quote: "quote",
			},
			Mock: struct {
				CreateQuote struct {
					Input  models.CreateQuoteModel
					Output models.QuoteModel
					Error  error
				}
			}{
				CreateQuote: struct {
					Input  models.CreateQuoteModel
					Output models.QuoteModel
					Error  error
				}{
					Input: models.CreateQuoteModel{
						ID:         id,
						Quote:      "quote",
						CreateDate: date,
						UpdateDate: date,
					},
					Output: models.QuoteModel{},
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
			quoteRepo.On("CreateQuote", mock.Anything).Return(c.Mock.CreateQuote.Output, c.Mock.CreateQuote.Error)

			quoteService := services.NewQuoteService(quoteRepo)
			result := quoteService.CreateQuote(c.Input.Quote)

			assert.Equal(t, c.Output, result)
		})
	}
}

func Test_UpdateQuote(t *testing.T) {
	type test struct {
		Name  string
		Input struct {
			ID    string
			Quote string
			Vote  int
		}
		Mock struct {
			UpdateQuote struct {
				Input struct {
					ID      string
					Payload models.UpdateQuoteModel
				}
				Output models.QuoteModel
				Error  error
			}
		}
		Output models.ResponseModel
	}
	id := uuid.New().String()
	quoteID := uuid.New().String()
	date := time.Now()
	cases := []test{
		{
			Name: "update quote success",
			Input: struct {
				ID    string
				Quote string
				Vote  int
			}{
				ID:    id,
				Quote: quoteID,
				Vote:  1,
			},
			Mock: struct {
				UpdateQuote struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}
			}{
				UpdateQuote: struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}{
					Input: struct {
						ID      string
						Payload models.UpdateQuoteModel
					}{
						ID: id,
						Payload: models.UpdateQuoteModel{
							Quote: quoteID,
							Vote:  1,
						},
					},
					Output: models.QuoteModel{
						ID:         id,
						Quote:      quoteID,
						UpdateDate: date,
					},
					Error: nil,
				},
			},
			Output: models.ResponseModel{
				Status:  true,
				Code:    200,
				Message: "update quote success",
				Result: models.QuoteModel{
					ID:         id,
					Quote:      quoteID,
					UpdateDate: date,
				},
			},
		},
		{
			Name: "id not found",
			Input: struct {
				ID    string
				Quote string
				Vote  int
			}{
				ID:    "",
				Quote: quoteID,
				Vote:  1,
			},
			Mock: struct {
				UpdateQuote struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}
			}{
				UpdateQuote: struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "id or quote not found",
				Result:  nil,
			},
		},
		{
			Name: "quote id not found",
			Input: struct {
				ID    string
				Quote string
				Vote  int
			}{
				ID:    id,
				Quote: "",
				Vote:  1,
			},
			Mock: struct {
				UpdateQuote struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}
			}{
				UpdateQuote: struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "id or quote not found",
				Result:  nil,
			},
		},
		{
			Name: "vote must be >= 0",
			Input: struct {
				ID    string
				Quote string
				Vote  int
			}{
				ID:    id,
				Quote: quoteID,
				Vote:  -1,
			},
			Mock: struct {
				UpdateQuote struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}
			}{
				UpdateQuote: struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "vote must be >= 0",
				Result:  nil,
			},
		},
		{
			Name: "update quote error",
			Input: struct {
				ID    string
				Quote string
				Vote  int
			}{
				ID:    id,
				Quote: quoteID,
				Vote:  1,
			},
			Mock: struct {
				UpdateQuote struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}
			}{
				UpdateQuote: struct {
					Input struct {
						ID      string
						Payload models.UpdateQuoteModel
					}
					Output models.QuoteModel
					Error  error
				}{
					Input: struct {
						ID      string
						Payload models.UpdateQuoteModel
					}{
						ID: id,
						Payload: models.UpdateQuoteModel{
							Quote: quoteID,
							Vote:  1,
						},
					},
					Output: models.QuoteModel{},
					Error:  errors.New("update quote error"),
				},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "update quote error",
				Result:  nil,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			quoteRepo := repositories.NewQuoteRepositoryMock()
			quoteRepo.On("UpdateQuote", mock.Anything, mock.Anything).Return(c.Mock.UpdateQuote.Output, c.Mock.UpdateQuote.Error)

			quoteService := services.NewQuoteService(quoteRepo)
			result := quoteService.UpdateQuote(c.Input.ID, c.Input.Quote, c.Input.Vote)

			assert.Equal(t, c.Output, result)
		})
	}
}
