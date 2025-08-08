package services

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rijwanansari/vivaLearning/config"
	"github.com/rijwanansari/vivaLearning/types"
	"github.com/rijwanansari/vivaLearning/utils/errutil"
	"github.com/vivasoft-ltd/golang-course-utils/logger"
)

type TokenServiceImpl struct {
	RedisService *RedisService
}

func NewTokenService(redisService *RedisService) *TokenServiceImpl {
	return &TokenServiceImpl{
		RedisService: redisService,
	}
}

// create a new token
func (s *TokenServiceImpl) CreateToken(userId int) (*types.Token, error) {
	jwtConf := config.Jwt()
	token := &types.Token{}

	token.UserID = userId
	token.AccessExpiry = time.Now().Add(jwtConf.GetAccessTokenExpiry()).Unix()
	token.RefreshExpiry = time.Now().Add(jwtConf.GetRefreshTokenExpiry()).Unix()
	token.AccessUuid = uuid.New().String()
	token.RefreshUuid = uuid.New().String()

	atClaims := jwt.MapClaims{}
	atClaims["uid"] = userId
	atClaims["aid"] = token.AccessUuid
	atClaims["exp"] = token.AccessExpiry
	atClaims["rid"] = token.RefreshUuid

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	var err error
	token.AccessToken, err = at.SignedString([]byte(jwtConf.AccessTokenSecret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["uid"] = userId
	rtClaims["aid"] = token.AccessUuid
	rtClaims["rid"] = token.RefreshUuid
	rtClaims["exp"] = token.RefreshExpiry

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(jwtConf.RefreshTokenSecret))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return token, nil

}

func (svc *TokenServiceImpl) StoreTokenUUID(token *types.Token) error {
	accessTokenCacheKey := config.Redis().MandatoryPrefix + config.Redis().AccessUuidPrefix + token.AccessUuid

	err := svc.RedisService.Set(accessTokenCacheKey, token.UserID, time.Duration(token.AccessExpiry))
	if err != nil {
		return err
	}

	refreshTokenCacheKey := config.Redis().MandatoryPrefix + config.Redis().RefreshUuidPrefix + token.RefreshUuid

	err = svc.RedisService.Set(refreshTokenCacheKey, token.UserID, time.Duration(token.RefreshExpiry))
	if err != nil {
		return err
	}

	return nil
}

func (svc *TokenServiceImpl) ParseAccessToken(accessToken string) (*types.Token, error) {
	parsedToken, err := ParseJwtToken(accessToken, config.Jwt().AccessTokenSecret)
	if err != nil {
		return nil, errutil.ErrParseJwt
	}

	if _, ok := parsedToken.Claims.(jwt.Claims); !ok || !parsedToken.Valid {
		return nil, errutil.ErrInvalidAccessToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errutil.ErrInvalidAccessToken
	}

	return mapClaimsToToken(claims)
}

func (svc *TokenServiceImpl) ReadUserIDFromAccessTokenUUID(accessTokenUuid string) (int, error) {
	accessTokenCacheKey := config.Redis().MandatoryPrefix + config.Redis().AccessUuidPrefix + accessTokenUuid

	userID, err := svc.RedisService.GetInt(accessTokenCacheKey)

	if err != nil {
		return 0, err
	}
	return userID, nil
}

func ParseJwtToken(token, secret string) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errutil.ErrInvalidJwtSigningMethod
		}
		return []byte(secret), nil
	}

	return jwt.Parse(token, keyFunc)
}

func mapClaimsToToken(claims jwt.MapClaims) (*types.Token, error) {
	jsonData, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	var token types.Token
	err = json.Unmarshal(jsonData, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
