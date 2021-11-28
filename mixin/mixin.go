package mixin

import (
	"encoding/hex"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
	"strconv"
	"time"
)

// PrivateKey 128
const PrivateKey = "cf67f23acb8f6f973811b5d99766639d8b28e2328d4a079a9fce2b6c95cc2041e7282cf923455db3cc2840369b133c9c3e6668b481e86397982023499aebea21"
// PublicKey 64
const PublicKey = "e7282cf923455db3cc2840369b133c9c3e6668b481e86397982023499aebea21"

func GetToken(user int) (string, error) {
	v2 := paseto.NewV2()
	b, _ := hex.DecodeString(PrivateKey)
	privateKey := ed25519.PrivateKey(b)
	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(24 * time.Hour),
		Issuer:     strconv.Itoa(user),
		IssuedAt:   time.Now(),
	}
	footer := strconv.Itoa(user)

	token, err := v2.Sign(privateKey, jsonToken, footer)
	return token, err
}
func VerifyToken(token string) (bool, paseto.JSONToken, string) {
	v2 := paseto.NewV2()
	b, _ := hex.DecodeString(PublicKey)
	publicKEY := ed25519.PublicKey(b)
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := v2.Verify(token, publicKEY, &newJsonToken, &newFooter)
	return err == nil, newJsonToken, newFooter
}