package middleware

import (
	"github.com/gofiber/fiber"
	"github.com/pascaloseko/shopping_list/server/util"
	"strings"
)

// Authz validates token and authorizes users
func Authz() func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		clientToken := ctx.Get("Authorization")
		if clientToken == "" {
			ctx.Status(403).JSON(&fiber.Map{
				"success": false,
				"message": "No Authorization header provided",
			})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			ctx.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": "Incorrect Format of Authorization Token",
			})
			return
		}

		jwtWrapper := util.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)

		if err != nil {
			ctx.Status(401).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		if claims != nil && claims.Email != "" {
			ctx.Locals("email", claims.Email)
			ctx.Next()
		}
	}
}
