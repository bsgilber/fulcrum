package application

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bsgilber/fulcrum/internal/controllers"
	"github.com/bsgilber/fulcrum/internal/forms"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type App struct {
	Srv *gin.Engine
}

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// RequestIDMiddleware ...
// Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

var auth = new(controllers.AuthController)

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func New() *App {
	// environment := initEnv()
	// cfg := config.NewConfig(environment, "./config")

	application := App{}

	// redis := databases.GetRedisClient(cfg.Redis.Host, cfg.Redis.Port)

	// psql := databases.GetPsqlClient(cfg.Psql.Uri)
	// h := controllers.NewBaseHandler(controllerDb, progenitorDb)

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	router.Use(CORSMiddleware())
	router.Use(RequestIDMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/health", health)
	application.Srv = router

	return &application
}

func (a *App) Start() {
	a.Srv.Run(":8080")
}

func initEnv() string {
	val, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		val = "local"
	}

	return val
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
