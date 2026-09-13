package engineRouter

import (
	"ddd-structure/internal/interface/http/middleware"
	"ddd-structure/internal/interface/http/routers"
	usersV1 "ddd-structure/internal/interface/http/routers/v1/users"

	"github.com/gin-gonic/gin"
)

func NewRouter(deps *routers.Dependencies) *gin.Engine {
	gin.SetMode(deps.Config.HTTPMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Private API (same idea as /api/private/v1 in szpt_new)
	api := r.Group("/api/private/v1")
	api.Use(middleware.DemoAuth())

	usersV1.SetupUsersV1Router(api, deps)

	// Add more modules here:
	// organizationsV1.SetupOrganizationsV1Router(api, deps)
	// documentsV1.SetupDocumentsV1Router(api, deps)

	return r
}
