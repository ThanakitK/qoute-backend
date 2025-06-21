package models

import "time"

type HandCreateQuoteBodyModel struct {
	Quote string `json:"quote"`
}

type RepoCreateQuoteModel struct {
	ID         string    `json:"id" bson:"id"`
	Quote      string    `json:"quote" bson:"quote"`
	Vote       int       `json:"vote" bson:"vote"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}

type RepoUpdateQuoteModel struct {
	Quote      string    `json:"quote" bson:"quote,omitempty"`
	Vote       int       `json:"vote" bson:"vote,omitempty"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}

type QuoteModel struct {
	ID         string    `json:"id" bson:"id"`
	Quote      string    `json:"quote" bson:"quote"`
	Vote       int       `json:"vote" bson:"vote"`
	CreateDate time.Time `json:"create_date" bson:"create_date"`
	UpdateDate time.Time `json:"update_date" bson:"update_date"`
}

type HandUpdateQuoteBodyModel struct {
	Quote string `json:"quote"`
	Vote  int    `json:"vote"`
}

type ResponseModel struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
