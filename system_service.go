package kaltura

type SystemService service

// Ping the server
func (s *SystemService) Ping() (bool, *ClientResponse, error) {
	u, err := buildQuery("system", "ping", nil)
	if err != nil {
		return false, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return false, nil, err
	}

	// execute action.
	var result = false
	res, err := s.client.Do(req, &result)
	if err != nil {
		return result, nil, err
	}

	return result, res, nil
}

func (s *SystemService) PingDatabase() (bool, *ClientResponse, error) {
	u, err := buildQuery("system", "pingDatabase", nil)
	if err != nil {
		return false, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return false, nil, err
	}

	// execute action.
	var result = false
	res, err := s.client.Do(req, &result)
	if err != nil {
		return result, nil, err
	}

	return result, res, nil
}

// Get server API version

func (s *SystemService) GetVersion() (string, *ClientResponse, error) {
	u, err := buildQuery("system", "getVersion", nil)
	if err != nil {
		return "", nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", nil, err
	}

	// execute action.
	var result = ""
	res, err := s.client.Do(req, &result)
	if err != nil {
		return result, nil, err
	}

	return result, res, nil
}

func (s *SystemService) GetTime() (int64, *ClientResponse, error) {
	u, err := buildQuery("system", "getTime", nil)
	if err != nil {
		return 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return 0, nil, err
	}

	// execute action.
	var result = int64(0)
	res, err := s.client.Do(req, &result)
	if err != nil {
		return result, nil, err
	}

	return result, res, nil
}

func (s *SystemService) GetHealthCheck() (string, *ClientResponse, error) {
	u, err := buildQuery("system", "GetHealthCheck", nil)
	if err != nil {
		return "", nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", nil, err
	}

	// execute action.
	var result = ""
	res, err := s.client.Do(req, nil)
	if err != nil {
		return result, nil, err
	}

	return result, res, nil
}
