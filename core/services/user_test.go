package services_test

import (
	"backend/core/models"
	"backend/core/repositories"
	"backend/core/services"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_SignIn(t *testing.T) {
	type test struct {
		Name  string
		Input struct {
			Email    string
			Password string
		}
		Mock struct {
			GetUser struct {
				Input  string
				Output models.UserModel
				Error  error
			}
		}
		Output models.ResponseModel
	}
	id := uuid.New().String()
	cases := []test{
		{
			Name: "success",
			Input: struct {
				Email    string
				Password string
			}{
				Email:    "test@gmail.com",
				Password: "123",
			},
			Mock: struct {
				GetUser struct {
					Input  string
					Output models.UserModel
					Error  error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{
					Input: "test@gmail.com",
					Output: models.UserModel{
						ID:         id,
						Email:      "test@gmail.com",
						Password:   "$2a$10$TODe5QSVwJdjrhPnpKPZb.uRL7dMA3YnOx6VCXcZs5HiPoYHs7c.6",
						QouteID:    "",
						CreateDate: time.Now(),
						UpdateDate: time.Now(),
					},
					Error: nil,
				},
			},
			Output: models.ResponseModel{
				Status:  true,
				Code:    200,
				Message: "sign in success",
				Result:  models.SignInResModel{},
			},
		},
		{
			Name: "email not found",
			Input: struct {
				Email    string
				Password string
			}{
				Email:    "",
				Password: "123",
			},
			Mock: struct {
				GetUser struct {
					Input  string
					Output models.UserModel
					Error  error
				}
			}{},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "email or password not found",
				Result:  nil,
			},
		},
		{
			Name: "password not found",
			Input: struct {
				Email    string
				Password string
			}{
				Email:    "test@gmail.com",
				Password: "",
			},
			Mock: struct {
				GetUser struct {
					Input  string
					Output models.UserModel
					Error  error
				}
			}{},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "email or password not found",
				Result:  nil,
			},
		},
		{
			Name: "email invalid",
			Input: struct {
				Email    string
				Password string
			}{
				Email:    "test",
				Password: "123",
			},
			Mock: struct {
				GetUser struct {
					Input  string
					Output models.UserModel
					Error  error
				}
			}{},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "email invalid",
				Result:  nil,
			},
		},
		{
			Name: "user is exist",
			Input: struct {
				Email    string
				Password string
			}{
				Email:    "test@gmail.com",
				Password: "123",
			},
			Mock: struct {
				GetUser struct {
					Input  string
					Output models.UserModel
					Error  error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{
					Input:  "test@gmail.com",
					Output: models.UserModel{},
					Error:  mongo.ErrNoDocuments,
				},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: mongo.ErrNoDocuments.Error(),
				Result:  nil,
			},
		},
		{
			Name: "password invalid",
			Input: struct {
				Email    string
				Password string
			}{
				Email:    "test@gmail.com",
				Password: "12345",
			},
			Mock: struct {
				GetUser struct {
					Input  string
					Output models.UserModel
					Error  error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{
					Input: "test@gmail.com",
					Output: models.UserModel{
						ID:         id,
						Email:      "test@gmail.com",
						Password:   "$2a$10$TODe5QSVwJdjrhPnpKPZb.uRL7dMA3YnOx6VCXcZs5HiPoYHs7c.6",
						QouteID:    "",
						CreateDate: time.Now(),
						UpdateDate: time.Now(),
					},
					Error: nil,
				},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "password invalid",
				Result:  nil,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("GetUser", c.Mock.GetUser.Input).Return(c.Mock.GetUser.Output, c.Mock.GetUser.Error)
			userService := services.NewUserService(userRepo)
			result := userService.SignIn(c.Input.Email, c.Input.Password)

			assert.Equal(t, result.Message, c.Output.Message)
		})
	}
}
