package common

type Authorization interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}

type auth struct {
	secretKey string
}
