package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/flyingpot/funcaptcha"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        token, err := funcaptcha.GetOpenAIToken()
        if err != nil {
            return c.JSON(fiber.Map{ "error": err.Error(), })
        }
        return c.JSON(fiber.Map{ "token": token, })
    })

    app.Listen(":8080")
}
