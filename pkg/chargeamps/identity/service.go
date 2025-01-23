package identity

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"

	"gitlab.com/gridio/test-assignment/internal"
	httphelper "gitlab.com/gridio/test-assignment/pkg/chargeamps/http"
)

type TokenSource struct {
	token internalToken
}

type internalToken struct {
	Message      string `json:"message"`
	Token        string `json:"token"`
	User         user   `json:"user"`
	RefreshToken string `json:"refreshToken"`
}

type user struct {
	Id         string     `json:"id"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	Mobile     string     `json:"mobile"`
	RfidTags   []struct{} `json:"rfidTags"`
	UserStatus string     `json:"userStatus"`
}

func CreateFromSecretAgent(_ logrus.FieldLogger, sa internal.SecretAgent) (*TokenSource, error) {
	t := TokenSource{}

	var unmarshalled string

	err := json.Unmarshal([]byte(sa.ProvideSecret()), &unmarshalled)
	if err != nil {
		return nil, err
	}

	t.token.Token = unmarshalled

	return &t, nil
}

func Login(_ logrus.FieldLogger, email string, password string) (*TokenSource, error) {
	const url = "https://eapi.charge.space/api/v5/auth/login"
	apiKey, _ := os.LookupEnv("APIKEY")
	postBody, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	responseBody := bytes.NewBuffer(postBody)

	req, err := httphelper.CreateRequest("POST", url, "apiKey", apiKey, responseBody)
	if err != nil {
		return nil, err
	}

	body, err := httphelper.GetResponseBody(req)
	if err != nil {
		return nil, err
	}

	var token internalToken
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	t := TokenSource{token: token}

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
	b, _ := json.Marshal(t.token.Token)

	return string(b)
}
