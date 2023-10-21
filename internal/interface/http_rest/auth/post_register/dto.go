package post_register

import "authsvc/internal/application/auth/register"

type requestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *requestBody) toUsecasePayload() *register.Payload {
	return &register.Payload{
		Username: r.Username,
		Password: r.Password,
	}
}
