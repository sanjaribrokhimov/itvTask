package routes

import (
	"net/http"
	"task_itv/config"
	"task_itv/handlers"
	"task_itv/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	engine       *gin.Engine
	config       *config.Config
	movieHandler *handlers.MovieHandler
	auth         *middleware.AuthMiddleware
}

func NewRouter(
	cfg *config.Config,
	movieHandler *handlers.MovieHandler,
	auth *middleware.AuthMiddleware,
) *Router {
	router := &Router{
		engine:       gin.Default(),
		config:       cfg,
		movieHandler: movieHandler,
		auth:         auth,
	}
	return router
}

func (r *Router) SetupRoutes() {
	// Swagger documentation
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.engine.POST("/login", r.login)

	// Protected routes with JWT authentication
	protected := r.engine.Group("/api")
	protected.Use(r.auth.Auth())
	{
		// Movie routes
		movies := protected.Group("/movies")
		{
			movies.POST("", r.movieHandler.CreateMovie)
			movies.GET("", r.movieHandler.GetMovies)
			movies.GET("/:id", r.movieHandler.GetMovie)
			movies.PUT("/:id", r.movieHandler.UpdateMovie)
			movies.DELETE("/:id", r.movieHandler.DeleteMovie)
		}
	}

	// Start server
	r.engine.Run(":" + r.config.Port)
}

func (r *Router) login(c *gin.Context) {
	// TODO: Implement proper login logic
	// For now, just return a valid JWT token
	token, err := r.auth.GenerateToken(1) // Using user ID 1 for now
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
