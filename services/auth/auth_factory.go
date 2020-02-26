package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/patrickmn/go-cache"
)

type AuthStrategy interface {
	GenerateToken(resource string) string
	ValidateToken(token string) (*jwt.Token, error)
}

func GetAuthStrategyService(cacheClient *cache.Cache) (AuthStrategy, error) {
	switch config.AUTH_STRATEGY {
	case "SIMPLE_JWT":
		return new(SimpleJwtAuth), nil
	case "TLS_CERTIFICATE_JWT":
		tlsProvider := new(TlsJwtAuth)
		tlsProvider.SetCacheProvider(cacheClient)
		return tlsProvider, nil
	default:
		return nil, errors.New("Invalid AuthStrategy Type")
	}
}
