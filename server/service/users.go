package service

type UsersResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type UsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersService interface {
	UserLogin(request UsersRequest) (*UsersResponse, error)
	UserSignUp(request UsersRequest) (*UsersResponse, error)
}
