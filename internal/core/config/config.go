package config

import "time"

type Auth struct {
	AccessTokenSecret, RefreshTokenSecret string
	AccessTokenTTL, RefreshTokenTTL       time.Duration
}

type Config struct {
	HttpPort           int
	PGConnectionString string
	Auth
}
