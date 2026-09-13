package validation

// UserValidation is the authenticated principal stored in gin context under key "user".
// Shape mirrors typical OIDC userinfo claims used in szpt_new.
type UserValidation struct {
	Sub        string `json:"sub"`
	Email      string `json:"email"`
	UserName   string `json:"preferred_username"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Reason string `json:"reason,omitempty"`
}
