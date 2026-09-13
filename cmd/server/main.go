package main

import (
	"fmt"
	"log"

	"ddd-structure/configs"
	usersApp "ddd-structure/internal/application/v1/users_use_case"
	usersPg "ddd-structure/internal/infrastructure/persistence/pg/v1/users"
	usersHandler "ddd-structure/internal/interface/http/handlers/v1/users"
	"ddd-structure/internal/interface/http/routers"
	"ddd-structure/internal/interface/http/routers/engineRouter"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Composition root: wire config → db → repo → use case → handler → router.
// Copy this pattern when adding a new module (see README).
func main() {
	cfg := configs.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("failed to connect to PostgreSQL:", err)
	}
	configs.DB = db

	// ---- users v1 (example module) ----
	userRepo := usersPg.NewUserRepository(db)
	userService := usersApp.NewService(userRepo)
	userHandlerInst := usersHandler.NewHandler(userService)

	router := engineRouter.NewRouter(&routers.Dependencies{
		UsersV1Handler: userHandlerInst,
		Config:         cfg,
	})

	addr := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Printf("listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
