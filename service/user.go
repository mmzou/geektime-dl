package service

//User user info
type User struct {
	UID       int    `json:"uid"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Cellphone string `json:"cellphone"`
}

//User user info
func (s *Service) User() (*User, Error) {
	body, err := s.requestUser()

	if err != nil {
		return nil, err
	}

	defer body.Close()

	user := new(User)
	if err := handleJSONParse(body, &user); err != nil {
		return nil, err
	}

	return user, nil
}
