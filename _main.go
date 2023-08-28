package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

type Person struct {
	Name string
	LastName string
}



func main() {
    jordi := Person{
        Name: "Jordi",
        LastName: "Polla",
    }
		
    engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Person": jordi,
        })
    })

    log.Fatal(app.Listen(":3000"))
}