package service

type UsersResponse struct {
	Username string
	Message  string
}

type UsersRequest struct {
	Username string
	Password string
}

type UsersService interface {
	UserLogin(request UsersRequest) (*UsersResponse, error)
	UserSignUp(request UsersRequest) (*UsersResponse, error)
}
