package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pmohanj/url-shorten-fiber-redis/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateReamaining int           `json:"rate_limit"`
	XRateLimitRest  time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse the request body",
		})
	}

	// Check if provided url is an actula url
	if !goValidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Provided URL is not valid",
		})
	}

	// Check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "You can't hack the system :]",
		})
	}

	// erforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	return nil
}
