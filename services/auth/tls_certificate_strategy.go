package auth

import (
	"bytes"
	"crypto/md5"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lcoutinho/luizalabs-client-api/config"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

type TlsJwtAuth struct {
	memoryCacheClient *cache.Cache
}

func (sv *TlsJwtAuth) SetCacheProvider(cacheClient *cache.Cache) {
	sv.memoryCacheClient = cacheClient
}

func (sv *TlsJwtAuth) GenerateToken(resource string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(config.JWT_TIME_EXPIRE).Unix()
	claims["resource"] = resource
	token.Claims = claims
	signingKey, _ := sv.generateTls()
	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		panic(err)
	}

	finalToken := fmt.Sprintf("%s", tokenString)

	keyMemory := sv.getMD5Hash(finalToken)

	sv.memoryCacheClient.Set(keyMemory, string(signingKey), config.JWT_TIME_EXPIRE)

	return fmt.Sprintf("%s", tokenString)
}

func (sv *TlsJwtAuth) ValidateToken(token string) (*jwt.Token, error) {

	keyMemory := sv.getMD5Hash(token)
	signingKey, found := sv.memoryCacheClient.Get(keyMemory)

	if !found {
		return nil, errors.New("Invalid token")
	}

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey.(string)), nil
	})
}

func (sv *TlsJwtAuth) generateTls() ([]byte, []byte) {

	var (
		err         error
		privKey     *rsa.PrivateKey
		pubKey      *rsa.PublicKey
		pubKeyBytes []byte
	)

	privKey, err = rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating private key")
	}
	pubKey = &privKey.PublicKey

	var privPEMBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	privKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(privKeyPEMBuffer, privPEMBlock)

	signingKey := privKeyPEMBuffer.Bytes()

	pubKeyBytes, err = x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		log.Fatal("Error marshalling public key")
	}

	var pubPEMBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	pubKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(pubKeyPEMBuffer, pubPEMBlock)
	verificationKey := pubKeyPEMBuffer.Bytes()

	return signingKey, verificationKey
}

func (sv *TlsJwtAuth) getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
