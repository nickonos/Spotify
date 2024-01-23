package api

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/nickonos/Spotify/packages/broker"
	"github.com/nickonos/Spotify/packages/routes"
)

func (api *API) GetSong(c *fiber.Ctx) error {
	err := ValidateRequest(c.GetReqHeaders()["Authorization"], "")
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).SendString(err.Error())
	}

	req := routes.GetSongRequest{
		Name: c.Query("name"),
	}

	if req.Name == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   "missing name query parameter",
		})
	}

	var res broker.Response[routes.GetSongResponse]
	err = broker.Request(api.broker, req, &res)
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
		"data":    res.Data.Song,
	})
}

func (api *API) CreateSong(c *fiber.Ctx) error {
	err := ValidateRequest(c.GetReqHeaders()["Authorization"], "admin")
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).SendString(err.Error())
	}

	var req routes.CreateSongRequest
	err = json.Unmarshal(c.BodyRaw(), &req)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	var res broker.Response[routes.CreateSongResponse]
	err = broker.Request(api.broker, req, &res)
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
		"data":    res.Data.Song,
	})
}

func (api *API) GetAllSongs(c *fiber.Ctx) error {
	err := ValidateRequest(c.GetReqHeaders()["Authorization"], "")
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).SendString(err.Error())
	}
	var req routes.GetSongsRequest

	var res broker.Response[routes.GetSongsResponse]
	err = broker.Request(api.broker, req, &res)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    res.Data.Songs,
	})
}
