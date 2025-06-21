package models

import "time"

type HandGetUserBodyModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignInResModel struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
	ID          string `json:"id"`
	QouteID     string `json:"quote_id"`
}

type UserModel struct {
	ID         string    `json:"id" bson:"id"`
	Email      string    `json:"email" bson:"email"`
	Password   string    `json:"password" bson:"password"`
	QouteID    string    `json:"quote_id" bson:"quote_id"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}

type CreateUserModel struct {
	ID         string    `json:"id" bson:"id"`
	Email      string    `json:"email" bson:"email"`
	QouteID    string    `json:"quote_id" bson:"quote_id"`
	Password   string    `json:"password" bson:"password"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}

type HandCreateUserBodyModel struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type HandSignInBodyModel struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UpdateUserModel struct {
	Email      string    `json:"email" bson:"email,omitempty"`
	QuoteID    string    `json:"quote_id" bson:"quote_id,omitempty"`
	Password   string    `json:"password" bson:"password,omitempty"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}
