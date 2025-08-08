package domain

import "github.com/rijwanansari/vivaLearning/types"

type (
	TokenService interface {
		CreateToken(userID int) (*types.Token, error)
		StoreTokenUUID(token *types.Token) error
		ParseAccessToken(accessToken string) (*types.Token, error)
		ReadUserIDFromAccessTokenUUID(accessTokenUuid string) (int, error)
	}
)
