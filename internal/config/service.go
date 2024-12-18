package config

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/samber/lo"
)

func New(version string) (*Service, error) {
	var envValues envValues

	err := env.Parse(&envValues)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	s := &Service{
		envValues: envValues,
		version:   version,
	}

	return s, nil
}

type Service struct {
	envValues envValues
	version   string
}

type envValues struct {
	DatabaseHostname           string `env:"DATABASE_HOSTNAME,required"`
	DatabaseName               string `env:"DATABASE_NAME,required"`
	DatabasePassword           string `env:"DATABASE_PASSWORD"`
	DatabasePort               string `env:"DATABASE_PORT,required"`
	DatabaseSSLMode            bool   `env:"DATABASE_SSL_MODE"            envDefault:"true"`
	DatabaseUsername           string `env:"DATABASE_USERNAME"`
	DisableRegistration        bool   `env:"DISABLE_REGISTRATION"         envDefault:"false"`
	JWTSecret                  string `env:"JWT_SECRET,required"`
	PhotoFolders               string `env:"PHOTO_FOLDERS"`
	ProxyAuthEmailHeader       string `env:"PROXY_AUTH_EMAIL_HEADER"      envDefault:"Remote-Email"`
	ProxyAuthNameHeader        string `env:"PROXY_AUTH_NAME_HEADER"       envDefault:"Remote-Name"`
	RedisURL                   string `env:"REDIS_URL,required"`
	ReverseProxyAuthentication bool   `env:"REVERSE_PROXY_AUTHENTICATION" envDefault:"false"`
	SessionSecret              string `env:"SESSION_SECRET,required"`
}

func (s *Service) Version() string {
	return s.version
}

func (s *Service) PhotoFolders() []string {
	paths := strings.Split(s.envValues.PhotoFolders, ",")
	paths = lo.Filter(paths, func(path string, _ int) bool { return path != "" })

	return paths
}

func (s *Service) JWTSecret() string {
	return s.envValues.JWTSecret
}

func (s *Service) SessionSecret() string {
	return s.envValues.SessionSecret
}

func (s *Service) DatabaseURL() string {
	connString := "postgres://"

	if s.envValues.DatabaseUsername != "" {
		connString = fmt.Sprintf("%s%s", connString, s.envValues.DatabaseUsername)

		if s.envValues.DatabasePassword != "" {
			encodedPassword := url.QueryEscape(s.envValues.DatabasePassword)
			connString = fmt.Sprintf("%s:%s", connString, encodedPassword)
		}

		connString += "@"
	}

	sslMode := "disable"
	if s.envValues.DatabaseSSLMode {
		sslMode = "require"
	}

	connString = fmt.Sprintf(
		"%s%s:%s/%s?sslmode=%s",
		connString,
		s.envValues.DatabaseHostname,
		s.envValues.DatabasePort,
		s.envValues.DatabaseName,
		sslMode,
	)

	return connString
}

func (s *Service) DisableRegistration() bool {
	return s.envValues.DisableRegistration
}

func (s *Service) ReverseProxyAuthentication() bool {
	return s.envValues.ReverseProxyAuthentication
}

func (s *Service) ProxyAuthEmailHeader() string {
	return s.envValues.ProxyAuthEmailHeader
}

func (s *Service) ProxyAuthNameHeader() string {
	return s.envValues.ProxyAuthNameHeader
}

func (s *Service) RedisURL() string {
	return s.envValues.RedisURL
}
