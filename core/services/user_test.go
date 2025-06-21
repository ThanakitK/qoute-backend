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

func Test_CreateUser(t *testing.T) {
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
			CreateUser struct {
				Input models.CreateUserModel
				Error error
			}
		}
		Output models.ResponseModel
	}
	id := uuid.New().String()
	cases := []test{
		{
			Name: "create success",
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
				CreateUser struct {
					Input models.CreateUserModel
					Error error
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
				CreateUser: struct {
					Input models.CreateUserModel
					Error error
				}{
					Input: models.CreateUserModel{
						ID:       id,
						Email:    "test@gmail.com",
						Password: "$2a$10$TODe5QSVwJdjrhPnpKPZb.uRL7dMA3YnOx6VCXcZs5HiPoYHs7c.6",
					},
					Error: nil,
				},
			},
			Output: models.ResponseModel{
				Status:  true,
				Code:    201,
				Message: "create user success",
				Result:  nil,
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
				CreateUser struct {
					Input models.CreateUserModel
					Error error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{},
				CreateUser: struct {
					Input models.CreateUserModel
					Error error
				}{},
			},
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
				CreateUser struct {
					Input models.CreateUserModel
					Error error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{},
				CreateUser: struct {
					Input models.CreateUserModel
					Error error
				}{},
			},
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
				CreateUser struct {
					Input models.CreateUserModel
					Error error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{},
				CreateUser: struct {
					Input models.CreateUserModel
					Error error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "email invalid",
				Result:  nil,
			},
		},
		{
			Name: "email already exist",
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
				CreateUser struct {
					Input models.CreateUserModel
					Error error
				}
			}{
				GetUser: struct {
					Input  string
					Output models.UserModel
					Error  error
				}{
					Input:  "test@gmail.com",
					Output: models.UserModel{},
					Error:  nil,
				},
				CreateUser: struct {
					Input models.CreateUserModel
					Error error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "email already exist",
				Result:  nil,
			},
		},
		{
			Name: "create user error",
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
				CreateUser struct {
					Input models.CreateUserModel
					Error error
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
				CreateUser: struct {
					Input models.CreateUserModel
					Error error
				}{
					Input: models.CreateUserModel{
						Email:    "test@gmail.com",
						Password: "123",
					},
					Error: errors.New("create user error"),
				},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "create user error",
				Result:  nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("GetUser", c.Mock.GetUser.Input).Return(c.Mock.GetUser.Output, c.Mock.GetUser.Error)
			userRepo.On("CreateUser", mock.Anything).Return(c.Mock.CreateUser.Error)
			userService := services.NewUserService(userRepo)
			result := userService.CreateUser(c.Input.Email, c.Input.Password)

			assert.Equal(t, result, c.Output)
		})
	}
}

func Test_UpdateVote(t *testing.T) {
	type test struct {
		Name  string
		Input struct {
			ID      string
			QouteID string
		}
		Mock struct {
			UpdateUser struct {
				Input struct {
					ID      string
					Payload models.UpdateUserModel
				}
				Output models.UserModel
				Error  error
			}
		}
		Output models.ResponseModel
	}
	id := uuid.New().String()
	qouteID := uuid.New().String()
	updateDate := time.Now()
	cases := []test{
		{
			Name: "update vote success",
			Input: struct {
				ID      string
				QouteID string
			}{
				ID:      id,
				QouteID: qouteID,
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}{
					Input: struct {
						ID      string
						Payload models.UpdateUserModel
					}{
						ID: id,
						Payload: models.UpdateUserModel{
							QuoteID: qouteID,
						},
					},
					Output: models.UserModel{
						ID:         id,
						QouteID:    qouteID,
						UpdateDate: updateDate,
					},
					Error: nil,
				},
			},
			Output: models.ResponseModel{
				Status:  true,
				Code:    200,
				Message: "update vote success",
				Result: models.UserModel{
					ID:         id,
					QouteID:    qouteID,
					UpdateDate: updateDate,
				},
			},
		},
		{
			Name: "id not found",
			Input: struct {
				ID      string
				QouteID string
			}{
				ID:      "",
				QouteID: qouteID,
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "id or qouteID not found",
				Result:  nil,
			},
		},
		{
			Name: "qouteID not found",
			Input: struct {
				ID      string
				QouteID string
			}{
				ID:      id,
				QouteID: "",
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}{},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "id or qouteID not found",
				Result:  nil,
			},
		},
		{
			Name: "update user error",
			Input: struct {
				ID      string
				QouteID string
			}{
				ID:      id,
				QouteID: qouteID,
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.UpdateUserModel
					}
					Output models.UserModel
					Error  error
				}{
					Input: struct {
						ID      string
						Payload models.UpdateUserModel
					}{
						ID: id,
						Payload: models.UpdateUserModel{
							QuoteID: qouteID,
						},
					},
					Output: models.UserModel{},
					Error:  errors.New("update user error"),
				},
			},
			Output: models.ResponseModel{
				Status:  false,
				Code:    400,
				Message: "update user error",
				Result:  nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(c.Mock.UpdateUser.Output, c.Mock.UpdateUser.Error)
			userService := services.NewUserService(userRepo)
			result := userService.UpdateVote(c.Input.ID, c.Input.QouteID)
			assert.Equal(t, result, c.Output)
		})
	}
}
