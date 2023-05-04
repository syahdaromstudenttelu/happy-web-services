package middleware

import (
	"happy-api-service/config"
	"happy-api-service/model/web"
	"happy-api-service/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(fiberCtx *fiber.Ctx) error {
	configEnv, err := util.LoadConfig("../")
	jwtTokenString := string(fiberCtx.Request().Header.Cookie("token"))

	if jwtTokenString == "" {
		webResponse := web.WebResponse[string]{
			Code:   fiber.StatusUnauthorized,
			Status: "failed",
			Data:   "unauthorized",
		}

		fiberCtx.Response().Header.SetStatusCode(fiber.StatusUnauthorized)
		return fiberCtx.JSON(webResponse)
	}

	jwtClaim := &config.JwtClaim{}
	jwtToken, err := jwt.ParseWithClaims(jwtTokenString, jwtClaim, func(t *jwt.Token) (interface{}, error) {
		return []byte(configEnv.JwtSecretKey), nil
	})

	if err != nil || !jwtToken.Valid {
		webResponse := web.WebResponse[string]{
			Code:   fiber.StatusUnauthorized,
			Status: "failed",
			Data:   "unauthorized",
		}

		fiberCtx.Response().Header.SetStatusCode(fiber.StatusUnauthorized)
		return fiberCtx.JSON(webResponse)
	}

	return fiberCtx.Next()
}
