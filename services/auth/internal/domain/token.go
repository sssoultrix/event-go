package domain

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AccessTokenClaims struct {
	UserID string
	Role   string
	Exp    int64
	JTI    string
}
