package kaltura

type SessionService service

// SessionTypes
const (
	SESSION_TYPE_USER  = 0
	SESSION_TYPE_ADMIN = 2
)

type Session struct {
	Secret      string `json:"secret"`
	UserId      string `json:"userId"`
	PartnerId   int    `json:"partnerId"`
	Expiry      int    `json:"expiry"`
	Privileges  string `json:"privileges"`
	Type        int    `json:"type"`
	SessionType int    `json:"sessionType"`

	ImpersonatedPartnerId int `json:"impersonatedPartnerId"`

	Ks string `json:"ks"`
}

type WidgetSession struct {
	Ks         string `json:"ks"`
	ObjectType string `json:"objectType"`
	PartnerId  int    `json:"partnerId"`
	UserId     string `json:"userId"`
}

// Start a session with Kaltura's server.

func (s *SessionService) Start(session *Session) (*ClientResponse, error) {
	u, err := buildQuery("session", "start", nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, session)
	if err != nil {
		return nil, err
	}

	// execute action.
	var result = ""
	res, err := s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	session.Ks = result

	return res, nil
}

func (s *SessionService) Get() (*Session, *ClientResponse, error) {
	u, err := buildQuery("session", "get", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result = &Session{}
	res, err := s.client.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, res, nil
}
