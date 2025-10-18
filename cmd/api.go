package main

import (
	"log"
	"time"

	"github.com/Auxesia23/velarsyapi/internal/handlers"
	"github.com/Auxesia23/velarsyapi/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
)

type application struct {
	cfg            config
	ServiceHandler handlers.ServiceHandler
	WorkHandler    handlers.WorkHandler
	ProjectHandler handlers.Projecthandler
	ImageHandler   handlers.ImageHandler
	UserHandler    handlers.UserHandler
}

type config struct {
	name               string
	port               string
	readTimeout        time.Duration
	writeTimeout       time.Duration
	idleTimeout        time.Duration
	maxRequestBodySize int
}

func NewApplication(cfg config,
	serviceHandler handlers.ServiceHandler,
	workHandler handlers.WorkHandler,
	projectHandler handlers.Projecthandler,
	imageHandler handlers.ImageHandler,
	userHandler handlers.UserHandler,
) *application {
	return &application{
		cfg:            cfg,
		ServiceHandler: serviceHandler,
		WorkHandler:    workHandler,
		ProjectHandler: projectHandler,
		ImageHandler:   imageHandler,
		UserHandler:    userHandler,
	}
}

func (app *application) mount() fasthttp.RequestHandler {
	r := fiber.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	users := r.Group("/auth")
	{
		users.Post("/login", app.UserHandler.LoginUserHandler)
		users.Post("/register", middlewares.JWTAuthMiddleware, app.UserHandler.CreateUserHandler)
		users.Get("/users", middlewares.JWTAuthMiddleware, app.UserHandler.GetAllUsersHandler)
		users.Put("/users/:id", middlewares.JWTAuthMiddleware, app.UserHandler.UpdateUserHandler)
		users.Delete("/users/:id", middlewares.JWTAuthMiddleware, app.UserHandler.DeleteUserHandler)

	}

	services := r.Group("/services")
	{
		services.Get("/", app.ServiceHandler.GetAllServicesHandler)
		services.Post("/", middlewares.JWTAuthMiddleware, app.ServiceHandler.CreateServiceHandler)
		services.Put("/:id", middlewares.JWTAuthMiddleware, app.ServiceHandler.UpdateServiceHandler)
		services.Delete("/:id", middlewares.JWTAuthMiddleware, app.ServiceHandler.DeleteServiceHandler)
	}

	work := r.Group("/works")
	{
		work.Get("/", app.WorkHandler.GetAllWorkHandler)
		work.Get("/:work_slug", app.WorkHandler.GetSingleWorkHandler)

		work.Post("/", middlewares.JWTAuthMiddleware, app.WorkHandler.CreateWorkHandler)
		work.Put("/:work_id", middlewares.JWTAuthMiddleware, app.WorkHandler.UpdateWorkHandler)
		work.Delete("/:work_id", middlewares.JWTAuthMiddleware, app.WorkHandler.DeleteWorkHandler)

		work.Post("/:work_id/projects", middlewares.JWTAuthMiddleware, app.ProjectHandler.CreateProjectHandler)

	}
	project := r.Group("/projects")
	{
		project.Get("/:project_slug", app.ProjectHandler.GetSingleProjectHandler)

		project.Put("/:project_id", middlewares.JWTAuthMiddleware, app.ProjectHandler.UpdateProjectHandler)
		project.Delete("/:project_id", middlewares.JWTAuthMiddleware, app.ProjectHandler.DeleteProjectHandler)
		project.Post("/:project_id/images", middlewares.JWTAuthMiddleware, app.ImageHandler.CreateImageHandler)
	}
	image := r.Group("/images", middlewares.JWTAuthMiddleware)
	{
		image.Delete("/:image_id", app.ImageHandler.DeleteImageHandler)
	}

	return r.Handler()
}

func (app *application) run(r fasthttp.RequestHandler) error {
	srv := &fasthttp.Server{
		Handler:            r,
		Name:               app.cfg.name,
		ReadTimeout:        app.cfg.readTimeout,
		WriteTimeout:       app.cfg.writeTimeout,
		IdleTimeout:        app.cfg.idleTimeout,
		MaxRequestBodySize: app.cfg.maxRequestBodySize,
	}
	log.Printf("Starting %s on port %s", app.cfg.name, app.cfg.port)
	return srv.ListenAndServe(":" + app.cfg.port)
}
