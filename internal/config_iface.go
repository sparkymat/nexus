package internal

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
	DisableRegistration() bool
	ReverseProxyAuthentication() bool
	ProxyAuthEmailHeader() string
	ProxyAuthNameHeader() string
	Version() string
}
