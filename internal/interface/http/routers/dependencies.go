package routers

import (
	"ddd-structure/configs"
	usersV1Handler "ddd-structure/internal/interface/http/handlers/v1/users"
)

// Dependencies is the composition root bag passed into routers.
// Add a field per module handler — same pattern as szpt_new.
type Dependencies struct {
	UsersV1Handler *usersV1Handler.Handler
	Config         *configs.Config
}
