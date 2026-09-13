package middleware

import (
	"net/http"
	"strings"

	"ddd-structure/internal/interface/dto/validation"
	"ddd-structure/package/response"

	"github.com/gin-gonic/gin"
)

// DemoAuth is a template middleware.
// In production replace with IdP userinfo validation (see szpt_new ValidationToken).
// Accepts: Authorization: Bearer <anything> and puts a demo user into context.
func DemoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			response.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing bearer token")
			c.Abort()
			return
		}

		user := validation.UserValidation{
			Sub:        "00000000-0000-0000-0000-000000000001",
			Email:      "demo@example.com",
			UserName:   "demo@example.com",
			GivenName:  "Demo",
			FamilyName: "User",
		}
		c.Set("user", user)
		c.Next()
	}
}
