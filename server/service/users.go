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
	userLogin(request UsersRequest) (*UsersResponse, error)
	userSignUp(request UsersRequest) (*UsersResponse, error)
}
