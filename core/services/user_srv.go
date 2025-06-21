package services

import (
	"backend/core/models"
	"backend/core/repositories"
	"backend/utils"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	SignIn(email string, password string) (result models.ResponseModel)

	CreateUser(email string, password string) (result models.ResponseModel)

	UpdateVote(id string, qouteID string) (result models.ResponseModel)
}

type UserSrv struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserSrv{
		userRepo: userRepo,
	}
}

func (s *UserSrv) SignIn(email string, password string) (result models.ResponseModel) {
	if email == "" || password == "" {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "email or password not found",
			Result:  nil,
		}
	}

	if !utils.IsEmail(email) {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "email invalid",
			Result:  nil,
		}
	}

	user, err := s.userRepo.GetUser(email)
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}
	if !utils.ComparePassword(user.Password, password) {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "password invalid",
			Result:  nil,
		}
	}
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}
	}
	data := models.SignInResModel{
		Type:        "Bearer",
		AccessToken: token,
		ID:          user.ID,
		QouteID:     user.QouteID,
	}

	return models.ResponseModel{
		Status:  true,
		Code:    200,
		Message: "sign in success",
		Result:  data,
	}
}

func (s *UserSrv) CreateUser(email string, password string) (result models.ResponseModel) {
	if email == "" || password == "" {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "email or password not found",
			Result:  nil,
		}
	}

	if !utils.IsEmail(email) {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "email invalid",
			Result:  nil,
		}
	}

	_, err := s.userRepo.GetUser(email)
	if err == nil {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "email already exist",
			Result:  nil,
		}
	}
	payload := models.CreateUserModel{
		ID:         uuid.New().String(),
		Email:      email,
		QouteID:    "",
		Password:   utils.GeneratePassword(password),
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}
	err = s.userRepo.CreateUser(payload)
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
		Message: "create user success",
		Result:  nil,
	}
}

func (s *UserSrv) UpdateVote(id string, qouteID string) (result models.ResponseModel) {
	if id == "" || qouteID == "" {
		return models.ResponseModel{
			Status:  false,
			Code:    400,
			Message: "id or qouteID not found",
			Result:  nil,
		}
	}
	payload := models.UpdateUserModel{
		QuoteID:    qouteID,
		UpdateDate: time.Now(),
	}
	res, err := s.userRepo.UpdateUser(id, payload)
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
		Message: "update vote success",
		Result:  res,
	}
}
