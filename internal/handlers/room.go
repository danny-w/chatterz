package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	uuid := guuid.New().String()
	url := fmt.Sprintf("room/%s", uuid)
	return c.Redirect(url)
}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")

	if uuid == "" {
		return
	}

}
