//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package configuration

import "time"

type Configuration interface {
	Config() *Config
}

type (
	Config struct {
		Debug             bool          `mapstructure:"DEBUG"`
		DebugLogFrequency time.Duration `mapstructure:"DEBUG_LOG_FREQUENCY"`

		GCPProjectID string `mapstructure:"GCP_PROJECT_ID"`

		ServerPort string `mapstructure:"PORT"`

		DatabaseUser        string        `mapstructure:"DATABASE_USER"`
		DatabasePassword    string        `mapstructure:"DATABASE_PASSWORD"`
		DatabaseHost        string        `mapstructure:"DATABASE_HOST"`
		DatabasePort        int           `mapstructure:"DATABASE_PORT"`
		DatabaseName        string        `mapstructure:"DATABASE_DBNAME"`
		DatabaseMaxOpen     int           `mapstructure:"DATABASE_CONNECTION_MAX_OPEN"`
		DatabaseMaxIdle     int           `mapstructure:"DATABASE_CONNECTION_MAX_IDLE"`
		DatabaseMaxLifetime time.Duration `mapstructure:"DATABASE_CONNECTION_MAX_LIFETIME"`
	}
)
