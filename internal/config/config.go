package config

const (
	DefaultToken = ""
)

var (
	DefaultConfig = Config{
		Token: DefaultToken,
		User:  DefaultUser,
	}
)

type Config struct {
	Token string `yaml:"GithubToken"`
	User  User   `yaml:"User,omitempty"`
}

func New() Config {
	return DefaultConfig
}
