package tokenizer

import "time"

type Tokenizer interface {
	CreateToken(employeeId uint, companyId uint) (string, error)
	CheckExpiration(token string) (bool, error)
	DecryptToken(token string) (uint, uint, error)
}

func New(
	expiration time.Duration,
	signKey []byte,
) Tokenizer {
	return &tokenizerImpl{
		accessExpiration: expiration,
		signKey:          signKey,
	}
}
