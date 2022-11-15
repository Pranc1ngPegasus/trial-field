//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package configuration

type Configuration interface {
	Config() *Config
}

type (
	Config struct {
		Debug bool `mapstructure:"DEBUG"`

		GCPProjectID string `mapstructure:"GCP_PROJECT_ID"`

		ServerPort string `mapstructure:"PORT"`
	}
)
