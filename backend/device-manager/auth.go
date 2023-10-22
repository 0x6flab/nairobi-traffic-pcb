package devicemanager

type auth struct {
	id    string
	token string
}

type Auth interface {
	Identify(token string) (id string, err error)
}

func NewAuth(id, token string) Auth {
	return &auth{
		id:    id,
		token: token,
	}
}

func (a *auth) Identify(token string) (id string, err error) {
	if token == a.token {
		return a.id, nil
	}

	return "", ErrAuthentication
}
