package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"spaces.maxlloyd.no/draw"
	"spaces.maxlloyd.no/pb"
	"spaces.maxlloyd.no/ws"
)

func main() {
	fmt.Println("HellO")
	app := fiber.New()

	drawState := draw.DrawState{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/init", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/purge", func(c *fiber.Ctx) error {
		drawState.MeshSequences = make([]pb.MeshSequence, 0)

		response, err := json.Marshal(drawState)

		if err != nil {
			return err
		}

		return c.Send(response)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		var sequence pb.MeshSequence
		err := json.Unmarshal(c.Body(), &sequence)

		if err != nil {
			return err
		}

		drawState.MeshSequences = append(drawState.MeshSequences, sequence)

		return c.Send(c.Body())
	})

	// Get new drwaings
	app.Get("/poll", func(c *fiber.Ctx) error {
		response, err := json.Marshal(drawState)

		if err != nil {
			return err
		}

		return c.Send(response)
	})

	ws.SetupSocketEndpoints(app)

	app.Listen(":3000")
}
