package infrastructure

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/interfaces/controllers"
)

type Router struct {
	port         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	echo         *echo.Echo
}

func NewRouter(c *Config) *Router {
	r := &Router{
		port:         c.Server.Port,
		readTimeout:  c.Server.ReadTimeout,
		writeTimeout: c.Server.WriteTimeout,
		echo:         echo.New(),
	}

	r.setRouting()

	return r
}

func (r *Router) setRouting() {
	userController := controllers.NewUserController()

	r.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	r.echo.GET("/users", func(c echo.Context) error {
		return userController.GetList(c)
	})
	r.echo.GET("/users/:id", func(c echo.Context) error {
		return userController.Get(c)
	})
}

func (r Router) Run() {
	if err := r.echo.Start(fmt.Sprintf(":%s", r.port)); err != nil {
		panic(err)
	}
}
