package users

import (
	"ddd-structure/internal/interface/http/routers"

	"github.com/gin-gonic/gin"
)

func SetupUsersV1Router(api *gin.RouterGroup, deps *routers.Dependencies) {
	group := api.Group("/users")
	{
		group.POST("", deps.UsersV1Handler.CreateUser)
		group.GET("", deps.UsersV1Handler.ListUsers)
		group.GET("/:id", deps.UsersV1Handler.GetUser)
		group.PUT("/:id", deps.UsersV1Handler.UpdateUser)
		group.DELETE("/:id", deps.UsersV1Handler.DeleteUser)
	}
}
