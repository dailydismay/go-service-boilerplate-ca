package viperconfig

import (
	"authsvc/internal/core/config"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func New() (*config.Config, error) {
	if err := loadDotenvFile(); err != nil {
		return nil, err
	}
	setDefaults()

	config := &config.Config{
		HttpPort: viper.GetInt("port"),
		Auth: config.Auth{
			AccessTokenTTL:     time.Duration(viper.GetInt("auth.access.ttl")) * time.Second,
			RefreshTokenTTL:    time.Duration(viper.GetInt("auth.refresh.ttl")) * time.Second,
			AccessTokenSecret:  viper.GetString("auth.access.secret"),
			RefreshTokenSecret: viper.GetString("auth.refresh.secret"),
		},
		PGConnectionString: viper.GetString("postgres.connectionURI"),
	}
	return config, nil
}

func loadDotenvFile() error {
	configType := os.Getenv("ENV")
	if configType == "" {
		configType = "dev"
	}

	if configType == "dev" || configType == "" {
		err := godotenv.Load(fmt.Sprintf("config/.env.%s", configType))
		if err != nil {
			return errors.Wrap(err, "failed to load local dotenv file")
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}

func setDefaults() {
	viper.SetDefault("port", "3000")
	viper.SetDefault("auth.access.ttl", time.Minute*5)
	viper.SetDefault("auth.refresh.ttl", time.Hour*2)
	viper.SetDefault("auth.access.secret", "01234567890123456789012345678912")
	viper.SetDefault("auth.refresh.secret", "01234567890123456789012345678912")
}
