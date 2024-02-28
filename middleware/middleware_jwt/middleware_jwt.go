package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	Filter       func(c *fiber.Ctx) bool
	Unauthorized fiber.Handler
	Decode       func(c *fiber.Ctx) (*jwt.MapClaims, error)
	Secret       string
	Expiry       int64
}

var ConfigDefault = Config{
	Filter:       nil,
	Decode:       nil,
	Unauthorized: nil,
	Secret:       "a_very_weak_secret",
}

func configDefault(config ...Config) Config {

	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.Secret == "" {
		cfg.Secret = ConfigDefault.Secret
	}

	if cfg.Decode == nil {

		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {

			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("authorization header is required")
			}

			token, err := jwt.ParseWithClaims(
				authHeader[7:],
				&jwt.MapClaims{},
				func(token *jwt.Token) (interface{}, error) {

					return []byte(cfg.Secret), nil
				},
			)

			if err != nil {
				return nil, errors.New("error parsing token")
			}

			claims, ok := token.Claims.(*jwt.MapClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("invalid token")
			}

			return claims, nil
		}
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

func Encode(claims *jwt.MapClaims, secret string) (string, error) {

	claimsNow := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := claimsNow.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("error creating a token")
	}

	return signedToken, nil

}

func New(config Config) fiber.Handler {

	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {

		if cfg.Filter != nil && cfg.Filter(c) {

			return c.Next()
		}

		claims, err := cfg.Decode(c)

		if err == nil {
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}
