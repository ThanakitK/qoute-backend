package handlers

import (
	"backend/core/models"
	"backend/core/services"

	"github.com/gofiber/fiber/v2"
)

type quoteHand struct {
	quoteService services.QuoteService
}

func NewQuoteHandler(quoteService services.QuoteService) quoteHand {
	return quoteHand{
		quoteService: quoteService,
	}
}

func (h quoteHand) GetQuotes(c *fiber.Ctx) error {
	result := h.quoteService.GetQuotes()
	return c.Status(result.Code).JSON(result)
}

func (h quoteHand) CreateQuote(c *fiber.Ctx) error {
	body := models.HandCreateQuoteBodyModel{}
	c.BodyParser(&body)

	result := h.quoteService.CreateQuote(body.Quote)
	return c.Status(result.Code).JSON(result)
}

func (h quoteHand) UpdateQuote(c *fiber.Ctx) error {
	body := models.HandUpdateQuoteBodyModel{}
	c.BodyParser(&body)

	result := h.quoteService.UpdateQuote(c.Params("id"), body.Quote, body.Vote)
	return c.Status(result.Code).JSON(result)
}

func (h quoteHand) DeleteQuote(c *fiber.Ctx) error {
	result := h.quoteService.DeleteQuote(c.Params("id"))

	return c.Status(result.Code).JSON(result)
}
