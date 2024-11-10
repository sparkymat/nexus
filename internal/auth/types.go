package auth

type Config interface {
	JWTSecret() string
	ProxyAuthNameHeader() string
	ProxyAuthEmailHeader() string
}
