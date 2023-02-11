package config

const (
	DefaultName  = "user"
	DefaultEmail = "example@example.com"
)

var (
	DefaultUser = User{
		Name:  DefaultName,
		Email: DefaultEmail,
	}
)

type User struct {
	Name  string `yaml:"name,omitempty"`
	Email string `yaml:"email,omitempty"`
}

func NewUser() User {
	return DefaultUser
}
