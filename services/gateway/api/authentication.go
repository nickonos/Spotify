package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/routes"
)

func (api *API) Login(c *fiber.Ctx) error {
	req := routes.Login{
		Code: c.Query("code"),
	}

	if req.Code == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   "missing code query parameter",
		})
	}

	var res broker.Response[routes.LoginResponse]
	err := broker.Request(api.broker, req, &res)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if res.Err != "" {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   res.Err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    res.Data.JWT,
	})
}
