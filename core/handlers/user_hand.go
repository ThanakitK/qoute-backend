package handlers

import (
	"backend/core/models"
	"backend/core/services"

	"github.com/gofiber/fiber/v2"
)

type userHand struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHand {
	return userHand{
		userService: userService,
	}
}
func (h userHand) CreateUser(c *fiber.Ctx) error {
	body := models.HandCreateUserBodyModel{}
	c.BodyParser(&body)
	result := h.userService.CreateUser(body.Email, body.Password)
	return c.Status(result.Code).JSON(result)
}

func (h userHand) SignIn(c *fiber.Ctx) error {
	body := models.HandSignInBodyModel{}
	c.BodyParser(&body)

	result := h.userService.SignIn(body.Email, body.Password)

	return c.Status(result.Code).JSON(result)
}

func (h userHand) UpdateVote(c *fiber.Ctx) error {
	result := h.userService.UpdateVote(c.Params("id"), c.Params("qouteID"))
	return c.Status(result.Code).JSON(result)
}
