package token_test

import (
	"testing"

	"github.com/DELIGHT-LABS/citadel-token-example/token"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	service := token.NewTokenService("test", "49bd37f2afb2d0c71c7be5db6081144afd68e40d4a16b8b5108bbdffcd594cd3e6a76529c6fae76cdd51f2f13692fc2b1bd2cddfe29945415f64fa382796d41d", "1h")

	uuid := "uuid1"
	token, err := service.GenerateToken(uuid)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claim, err := service.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, uuid, claim.ID)
}
