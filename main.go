package main

import "github.com/gofiber/fiber/v2"

type SomeStruct struct {
	Name        string
	Ingredients string
}

func main() {
	app := fiber.New()

	app.Static("/", "./fiber-app/dist")

	app.Get("/api/v1/pizza", func(c *fiber.Ctx) error {
		data := SomeStruct{
			Name:        "Garden Party",
			Ingredients: "Nom nom",
		}

		return c.JSON(data)
	})

	_ = app.Listen(":3000")
}
