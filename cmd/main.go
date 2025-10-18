package main

import (
	"log"
	"time"

	"github.com/Auxesia23/velarsyapi/internal/database"
	"github.com/Auxesia23/velarsyapi/internal/handlers"
	"github.com/Auxesia23/velarsyapi/internal/repositories"
	"github.com/Auxesia23/velarsyapi/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	db, err := database.InitPostgres()
	if err != nil {

		log.Fatal(err)
	}
	cld, err := database.InitCloudinary()
	if err != nil {
		log.Fatal(err)
	}

	serviceRepository := repositories.NewServiceRepository(db)
	imageReposiry := repositories.NewImageRepository(cld, db)
	projectRepository := repositories.NewProjectRepository(db)
	workRepository := repositories.NewWorkRepository(db)
	userRepository := repositories.NewUserRepository(db)

	serviceService := services.NewServiceService(serviceRepository)
	projectService := services.NewProjectService(projectRepository, imageReposiry)
	workService := services.NewWorkService(workRepository, imageReposiry, projectRepository)
	imageServive := services.NewImageService(imageReposiry)
	userService := services.NewUserService(userRepository)

	workHandler := handlers.NewWorkHandler(workService)
	projectHandler := handlers.NewProjectHandler(projectService)
	serviveHandler := handlers.NewServiceHandler(serviceService)
	imageHandler := handlers.NewImageHandler(imageServive)
	userHandler := handlers.NewUserHandler(userService)

	cfg := config{
		name:         "Velarsy API",
		port:         "8080",
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
		idleTimeout:  10 * time.Second,
		maxRequestBodySize: 50 * 1024 * 1024,
	}
	app := NewApplication(cfg, serviveHandler, workHandler, projectHandler, imageHandler, userHandler)
	r := app.mount()
	log.Fatal(app.run(r))
}
