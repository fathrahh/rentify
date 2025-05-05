package route

import (
	"github.com/gofiber/fiber/v2"
	"ijor.dev/rentify/internal/delivery/http"
)

type RouteConfig struct {
	App         *fiber.App
	UserHandler *http.UserHandler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthenticatedRoute()
}

func (c *RouteConfig) SetupGuestRoute() {

}

func (c *RouteConfig) SetupAuthenticatedRoute() {

}
