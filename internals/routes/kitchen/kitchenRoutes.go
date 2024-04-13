package kitchenRoutes

import (
	"fmt"

	model "github.com/Object-Oriented-Project/backend/internals/model"
	"encoding/json"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupKitchenRoutes(route fiber.Router, app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/ws/kitchen" {
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)
				return c.Next()
			}
			return fiber.ErrUpgradeRequired
		}
		return c.Next()
	})
	clients := make(map[string]string)
	socketio.New(func(kws *socketio.Websocket) {
		userID := kws.Params("id")
		clients[userID] = kws.UUID
		kws.SetAttribute("user_id", userID)
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userID, kws.UUID)), true, socketio.TextMessage)
		// kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userID, kws.UUID)), socketio.TextMessage)
	})

    // socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
    //     fmt.Println(fmt.Sprintf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id")))
    // })
    // socketio.On("OrderItem", func(ep *socketio.EventPayload) {
    //     fmt.Println(fmt.Sprintf("Custom event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    // })
    // socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {

        // fmt.Println(fmt.Sprintf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data)))

        // message := model.Order{}

        // err := json.Unmarshal(ep.Data, &message)
        // if err != nil {
        //     fmt.Println(err)
        //     return
        // }
        // if message.CustomerName != "" {
        //     ep.Kws.Fire(message.CustomerName, []byte(message.Data))
        // }

        // err = ep.Kws.EmitTo(clients[message.To], ep.Data, socketio.TextMessage)
        // if err != nil {
        //     fmt.Println(err)
        // }
    // })
    socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        // fmt.Println(fmt.Sprintf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    })

    socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {
        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        // fmt.Println(fmt.Sprintf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    })

    // socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
    //     fmt.Println(fmt.Sprintf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
    // })
	route.Get("/ws/kitchen", socketio.New(func(kws *socketio.Websocket) {
		userId := kws.Params("id")
	
		clients[userId] = kws.UUID
	
		kws.SetAttribute("user_id", userId)
	
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true, socketio.TextMessage)
		// kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)), socketio.TextMessage)
	}))

	route.Post("/", 
		func(c *fiber.Ctx) error {
			order := new(model.Order)
    if err := c.BodyParser(order); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }
    orderBytes, err := json.Marshal(order)
    if err != nil {
        fmt.Println("Error marshalling order:", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
    }
    socketio.Broadcast(orderBytes, 1, socketio.TextMessage)
    return c.JSON(order)
		}, 
	)
}