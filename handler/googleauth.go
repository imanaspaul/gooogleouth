package handler

import (
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gofiber/fiber/v2"
)

func GoogleOauth(c *fiber.Ctx) error {

	v := googleAuthIDTokenVerifier.Verifier{}
	aud := "812743196970-ivek7rj733c9g2b7c2pbdrf9fta2q7bv.apps.googleusercontent.com"

	err := v.VerifyIDToken(c.Query("idtoken"), []string{aud})

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	claimSet, err := googleAuthIDTokenVerifier.Decode(c.Query("idtoken"))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": claimSet,
	})
}
