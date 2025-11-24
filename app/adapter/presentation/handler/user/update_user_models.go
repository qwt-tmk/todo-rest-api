package user

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserResponse struct {
	ID    string `json:"ID"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
