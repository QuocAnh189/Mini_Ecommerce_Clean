package token

const (
	AccessTokenType  = "x-access"  // 5 minutes
	RefreshTokenType = "x-refresh" // 30 days
)

type IMarker interface {
	GenerateAccessToken(payload *AuthPayload) string
	GenerateRefreshToken(payload *AuthPayload) string
	ValidateToken(token string) (*AuthPayload, error)
}
