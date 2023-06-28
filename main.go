package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/flyingpot/funcaptcha"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())

    app.Get("/token", func(c *fiber.Ctx) error {
        token, err := funcaptcha.GetOpenAIToken()
        if err != nil {
            return c.JSON(fiber.Map{ "error": err.Error(), })
        }
        return c.JSON(fiber.Map{ "token": token, })
    })

    app.Listen(":8080")
}
