package service

//User user info
func (s *Service) User() (*User, Error) {
	body, err := s.RequestUser()

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
