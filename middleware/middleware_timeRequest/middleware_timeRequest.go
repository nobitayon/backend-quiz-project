package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter       func(c *fiber.Ctx) bool
	Record       func(c *fiber.Ctx) (time.Time, error)
	Unauthorized func(c *fiber.Ctx) error
}

var ConfigDefault = Config{
	Filter:       nil,
	Record:       nil,
	Unauthorized: nil,
}

func configDefault(config ...Config) Config {

	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.Record == nil {

		cfg.Record = func(c *fiber.Ctx) (time.Time, error) {

			return time.Now(), nil
		}
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return cfg
}

func New(config Config) fiber.Handler {

	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {

		if cfg.Filter != nil && cfg.Filter(c) {

			return c.Next()
		}

		timeRecorded, err := cfg.Record(c)

		if err == nil {
			c.Locals("timeRecorded", timeRecorded)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}
