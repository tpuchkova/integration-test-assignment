package internal

type Agent struct {
	token string
}

func NewSecretAgent(token string) *Agent {
	a := Agent{
		token: token,
	}

	return &a
}

func (a *Agent) ProvideSecret() string {
	return a.token
}

func (a *Agent) UpdateSecret(s string) {
	a.token = s
}
