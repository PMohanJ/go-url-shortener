package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pmohanj/url-shorten-fiber-redis/database"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not URL found for given shorten url",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Can't connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}
