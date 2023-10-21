package post_login

import (
	"authsvc/internal/application/auth/login"
	"time"
)

type requestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *requestBody) toUsecasePayload() *login.Payload {
	return &login.Payload{
		Username: r.Username,
		Password: r.Password,
	}
}

type token struct {
	ID        string    `json:"id"`
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}

type responseBody struct {
	Access  *token `json:"access_token"`
	Refresh *token `json:"refreh_token"`
}

func responseFromResult(r *login.Result) *responseBody {
	return &responseBody{
		Access: &token{
			ID:        r.Access.ID,
			Value:     r.Access.Value,
			ExpiresAt: r.Access.ExpiresAt,
		},
		Refresh: &token{
			ID:        r.Refresh.ID,
			Value:     r.Refresh.Value,
			ExpiresAt: r.Refresh.ExpiresAt,
		},
	}
}
