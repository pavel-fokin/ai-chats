package api

import (
	"ai-chats/internal/domain"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(t *testing.T) {
	userID := domain.NewUserID()
	InitSigningKey("secret")
	tokenString, err := NewAccessToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Verify the token claims
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(*JWTClaims)
	assert.True(t, ok)
	assert.Equal(t, accessTokenIssuer, claims.Issuer)
	assert.Equal(t, accessTokenSubject, claims.Subject)
	assert.Equal(t, jwt.ClaimStrings{accessTokenAudience}, claims.Audience)
	assert.WithinDuration(t, time.Now().Add(accessTokenDuration), claims.ExpiresAt.Time, time.Second)
	assert.Equal(t, userID, claims.UserID)
}
