package identity

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"gitlab.com/gridio/test-assignment/internal"
)

type TokenSource struct {
	token internalToken
}

type internalToken struct {
	// TODO: Replace with proper fields that chargeamps is responding with
	Field1 string `json:"field1"`
}

func CreateFromSecretAgent(_ logrus.FieldLogger, sa internal.SecretAgent) *TokenSource {
	t := TokenSource{}

	var unmarshalled internalToken

	_ = json.Unmarshal([]byte(sa.ProvideSecret()), &unmarshalled)
	// TODO: Check error here

	t.token = unmarshalled

	return &t
}

func Login(_ logrus.FieldLogger, username string, password string) (*TokenSource, error) {
	t := TokenSource{}

	return &t, nil
}

func (t *TokenSource) AccessToken() string {
	// TODO implement me
	panic("implement me")
}

func (t *TokenSource) IsUnauthorized() bool {
	// TODO implement me
	panic("implement me")
}

func (t *TokenSource) String() string {
	b, _ := json.Marshal(t.token)

	return string(b)
}

// TODO: Write a function that retrieves access and refresh tokens from chargeamps and stores them in internalToken
// 	struct
